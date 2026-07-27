/*
Copyright (C) 2023-2026 Project Contributors
Modified for Aliyun Captcha 2.0 human verification (AGPLv3).

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
import { useCallback, useEffect, useRef } from 'react'

import { useStatus } from '@/hooks/use-status'

// Aliyun requires the captcha JS to be loaded dynamically from its CDN.
// Self-hosting or re-importing it breaks verification, so load it exactly once.
const CAPTCHA_SCRIPT_URL =
  'https://o.alicdn.com/captcha-frontend/aliyunCaptcha/AliyunCaptcha.js'

interface AliyunCaptchaInstance {
  show?: () => void
  hide?: () => void
}

interface AliyunCaptchaWindow extends Window {
  AliyunCaptchaConfig?: { region: string; prefix: string }
  initAliyunCaptcha?: (config: Record<string, unknown>) => void
}

// Memoize the script load globally so multiple hook instances (sign-up,
// sign-in, forgot-password) never import the captcha JS more than once.
let captchaScriptPromise: Promise<void> | null = null
function loadCaptchaScript(): Promise<void> {
  const w = window as AliyunCaptchaWindow
  if (typeof w.initAliyunCaptcha === 'function') return Promise.resolve()
  if (captchaScriptPromise) return captchaScriptPromise
  captchaScriptPromise = new Promise<void>((resolve, reject) => {
    const script = document.createElement('script')
    script.src = CAPTCHA_SCRIPT_URL
    script.async = true
    script.onload = () => resolve()
    script.onerror = () => {
      captchaScriptPromise = null
      reject(new Error('failed to load aliyun captcha script'))
    }
    document.head.appendChild(script)
  })
  return captchaScriptPromise
}

// Unique suffix generator so several mounted forms don't collide on DOM ids.
let captchaElementSeq = 0

/**
 * Hook for Aliyun Captcha 2.0 human verification (popup mode).
 *
 * Exposes getCaptchaVerifyParam(), which pops the captcha and resolves with the
 * captchaVerifyParam once the user passes verification, or rejects if the user
 * dismisses the popup / it times out. When captcha is disabled it resolves ''.
 * The returned param is single-use and must be forwarded verbatim to the
 * backend, which validates it via VerifyIntelligentCaptcha.
 */
export function useAliyunCaptcha() {
  const { status } = useStatus()
  const isEnabled = !!(
    status?.aliyun_captcha_enabled &&
    status?.aliyun_captcha_prefix &&
    status?.aliyun_captcha_scene_id
  )
  const prefix = (status?.aliyun_captcha_prefix as string) || ''
  const sceneId = (status?.aliyun_captcha_scene_id as string) || ''

  const instanceRef = useRef<AliyunCaptchaInstance | null>(null)
  const triggerRef = useRef<HTMLButtonElement | null>(null)
  const resolverRef = useRef<{
    resolve: (param: string) => void
    reject: (err: Error) => void
  } | null>(null)
  const readyRef = useRef(false)
  const idsRef = useRef<{ element: string; button: string } | null>(null)
  const initCaptchaRef = useRef<() => void>(() => {})

  const initCaptcha = useCallback(() => {
    const w = window as AliyunCaptchaWindow
    if (typeof w.initAliyunCaptcha !== 'function' || !idsRef.current) return
    readyRef.current = false
    w.initAliyunCaptcha({
      SceneId: sceneId,
      mode: 'popup',
      element: `#${idsRef.current.element}`,
      button: `#${idsRef.current.button}`,
      success: (captchaVerifyParam: string) => {
        const resolver = resolverRef.current
        resolverRef.current = null
        resolver?.resolve(captchaVerifyParam)
        // Aliyun params are single-use; re-initialize so the next request gets
        // a fresh captcha (documented re-verification flow).
        initCaptchaRef.current()
      },
      fail: () => {
        // Aliyun auto-refreshes the captcha within the valid period; no action.
      },
      getInstance: (instance: AliyunCaptchaInstance) => {
        instanceRef.current = instance
        readyRef.current = true
      },
      onClose: (reason: string) => {
        if (reason === 'userDismiss') {
          const resolver = resolverRef.current
          resolverRef.current = null
          resolver?.reject(new Error('captcha dismissed'))
        }
      },
      slideStyle: { width: 360, height: 40 },
      language: 'cn',
    })
  }, [sceneId])

  useEffect(() => {
    initCaptchaRef.current = initCaptcha
  }, [initCaptcha])

  useEffect(() => {
    if (!isEnabled) return
    let cancelled = false

    // Create the (hidden) captcha container + trigger button once. initAliyunCaptcha
    // requires both element and button selectors even though we trigger via show().
    if (!idsRef.current) {
      const seq = ++captchaElementSeq
      const elementId = `aliyun-captcha-element-${seq}`
      const buttonId = `aliyun-captcha-trigger-${seq}`
      const element = document.createElement('div')
      element.id = elementId
      document.body.appendChild(element)
      const button = document.createElement('button')
      button.id = buttonId
      button.type = 'button'
      button.style.display = 'none'
      document.body.appendChild(button)
      triggerRef.current = button
      idsRef.current = { element: elementId, button: buttonId }
    }

    ;(window as AliyunCaptchaWindow).AliyunCaptchaConfig = {
      region: 'cn',
      prefix,
    }

    loadCaptchaScript()
      .then(() => {
        if (!cancelled) initCaptcha()
      })
      .catch(() => {
        // CDN blocked; getCaptchaVerifyParam will time out and reject later.
      })

    return () => {
      cancelled = true
    }
  }, [isEnabled, prefix, initCaptcha])

  const getCaptchaVerifyParam = useCallback((): Promise<string> => {
    if (!isEnabled) return Promise.resolve('')
    return new Promise<string>((resolve, reject) => {
      resolverRef.current = { resolve, reject }
      const trigger = () => {
        if (instanceRef.current?.show) {
          instanceRef.current.show()
        } else {
          triggerRef.current?.click()
        }
      }
      if (readyRef.current) {
        trigger()
      } else {
        // Instance not ready yet; poll briefly until initialization completes.
        let tries = 0
        const timer = window.setInterval(() => {
          tries += 1
          if (readyRef.current) {
            window.clearInterval(timer)
            trigger()
          } else if (tries > 50) {
            window.clearInterval(timer)
            if (resolverRef.current) {
              resolverRef.current = null
              reject(new Error('captcha not ready'))
            }
          }
        }, 100)
      }
      // Safety timeout so callers never hang indefinitely.
      window.setTimeout(() => {
        if (resolverRef.current) {
          const resolver = resolverRef.current
          resolverRef.current = null
          resolver.reject(new Error('captcha timeout'))
        }
      }, 120000)
    })
  }, [isEnabled])

  return { isEnabled, getCaptchaVerifyParam }
}
