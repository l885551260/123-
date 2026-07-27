import { useState, useCallback, useRef } from 'react'
import { toast } from 'sonner'
import { useTranslation } from 'react-i18next'

import {
  requestAlipayNativePayment,
  requestWxPayNativePayment,
  isApiSuccess,
} from '../api'
import { isAlipayNativePayment } from '../lib'

export interface NativePaymentState {
  qrCode: string | null
  payUrl: string | null
  tradeNo: string | null
  processing: boolean
  polling: boolean
}

/**
 * Hook for native Alipay/WeChat Pay direct payment flow.
 * Handles QR code display (PC) and H5 redirect (mobile).
 */
export function useNativePayment(onSuccess?: () => void) {
  const { t } = useTranslation()
  const [state, setState] = useState<NativePaymentState>({
    qrCode: null,
    payUrl: null,
    tradeNo: null,
    processing: false,
    polling: false,
  })
  const pollTimerRef = useRef<ReturnType<typeof setInterval> | null>(null)

  const stopPolling = useCallback(() => {
    if (pollTimerRef.current) {
      clearInterval(pollTimerRef.current)
      pollTimerRef.current = null
    }
    setState((prev) => ({ ...prev, polling: false }))
  }, [])

  const startPolling = useCallback(() => {
    setState((prev) => ({ ...prev, polling: true }))
    let attempts = 0
    const maxAttempts = 100 // ~5 minutes at 3s interval

    pollTimerRef.current = setInterval(async () => {
      attempts++
      if (attempts > maxAttempts) {
        stopPolling()
        toast.error(t('Payment timeout, please check your order status'))
        return
      }

      try {
        const res = await fetch('/api/user/topup/self', {
          credentials: 'include',
        })
        const data = await res.json()
        if (data.success && data.data) {
          const orders = Array.isArray(data.data) ? data.data : []
          const latestOrder = orders[0]
          if (latestOrder && latestOrder.status === 'success') {
            stopPolling()
            setState((prev) => ({ ...prev, qrCode: null, payUrl: null, tradeNo: null }))
            toast.success(t('Payment successful!'))
            onSuccess?.()
          }
        }
      } catch {
        // Ignore polling errors
      }
    }, 3000)
  }, [stopPolling, t, onSuccess])

  const processNativePayment = useCallback(
    async (amount: number, paymentType: string) => {
      try {
        setState((prev) => ({ ...prev, processing: true }))

        const isAlipay = isAlipayNativePayment(paymentType)
        const response = isAlipay
          ? await requestAlipayNativePayment({ amount, payment_method: paymentType })
          : await requestWxPayNativePayment({ amount, payment_method: paymentType })

        if (!isApiSuccess(response)) {
          toast.error((response as { data?: string }).data as string || t('Payment request failed'))
          return false
        }

        const data = (response as { data?: { qr_code?: string; pay_url?: string; trade_no?: string } }).data
        if (!data) {
          toast.error(t('Payment request failed'))
          return false
        }

        if (data.qr_code) {
          // PC: Show QR code
          setState((prev) => ({
            ...prev,
            qrCode: data.qr_code!,
            tradeNo: data.trade_no || null,
            processing: false,
          }))
          startPolling()
          return true
        }

        if (data.pay_url) {
          // H5: Redirect to payment page
          setState((prev) => ({
            ...prev,
            payUrl: data.pay_url!,
            tradeNo: data.trade_no || null,
            processing: false,
          }))
          startPolling()
          window.location.href = data.pay_url
          return true
        }

        toast.error(t('Payment request failed'))
        return false
      } catch {
        toast.error(t('Payment request failed'))
        return false
      } finally {
        setState((prev) => ({ ...prev, processing: false }))
      }
    },
    [startPolling, stopPolling, t]
  )

  const reset = useCallback(() => {
    stopPolling()
    setState({
      qrCode: null,
      payUrl: null,
      tradeNo: null,
      processing: false,
      polling: false,
    })
  }, [stopPolling])

  return {
    ...state,
    processNativePayment,
    reset,
    stopPolling,
  }
}
