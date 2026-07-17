/*
Copyright (C) 2023-2026 QuantumNous
Modified by aytdai for phone/SMS registration feature (AGPLv3).

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.
*/
import i18next from 'i18next'
import { useState } from 'react'
import { toast } from 'sonner'

import { useCountdown } from '@/hooks/use-countdown'

import { sendSMSCode } from '../api'
import { SMS_VERIFICATION_COUNTDOWN } from '../constants'

interface UseSMSVerificationOptions {
  turnstileToken?: string
  validateTurnstile?: () => boolean
}

/**
 * Hook for managing SMS verification code sending
 */
export function useSMSVerification(options?: UseSMSVerificationOptions) {
  const [isSending, setIsSending] = useState(false)
  const {
    secondsLeft,
    isActive,
    start: startCountdown,
  } = useCountdown({ initialSeconds: SMS_VERIFICATION_COUNTDOWN })

  /**
   * Send verification code to phone
   */
  const sendCode = async (phone: string) => {
    if (!phone) {
      toast.error(i18next.t('Please enter your phone number'))
      return false
    }

    // Validate turnstile if validation function is provided
    if (options?.validateTurnstile && !options.validateTurnstile()) {
      return false
    }

    setIsSending(true)
    try {
      const res = await sendSMSCode(phone, options?.turnstileToken)
      if (res?.success) {
        startCountdown()
        toast.success(i18next.t('Verification code sent'))
        return true
      }
      toast.error(
        res?.message || i18next.t('Failed to send verification code')
      )
      return false
    } catch (_error) {
      // Errors are handled by global interceptor
      return false
    } finally {
      setIsSending(false)
    }
  }

  return {
    isSending,
    secondsLeft,
    isActive,
    sendCode,
  }
}
