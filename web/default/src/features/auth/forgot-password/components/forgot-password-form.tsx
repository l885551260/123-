/*
Copyright (C) 2023-2026 Project Contributors

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

For commercial licensing, please contact support@quantumnous.com
*/
import { zodResolver } from '@hookform/resolvers/zod'
import { ArrowRight, Loader2 } from 'lucide-react'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'
import type { z } from 'zod'

import { PasswordInput } from '@/components/password-input'
import { Turnstile } from '@/components/turnstile'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
  resetPasswordByPhone,
  sendPasswordResetEmail,
} from '@/features/auth/api'
import {
  forgotPasswordFormSchema,
  PASSWORD_RESET_COUNTDOWN,
  phoneResetFormSchema,
} from '@/features/auth/constants'
import { useAliyunCaptcha } from '@/features/auth/hooks/use-aliyun-captcha'
import { useSMSVerification } from '@/features/auth/hooks/use-sms-verification'
import { useTurnstile } from '@/features/auth/hooks/use-turnstile'
import { useCountdown } from '@/hooks/use-countdown'
import { useStatus } from '@/hooks/use-status'
import { cn } from '@/lib/utils'

type ForgotPasswordFormProps = React.HTMLAttributes<HTMLFormElement> & {
  resetMode: 'email' | 'phone'
  onResetModeChange: (mode: 'email' | 'phone') => void
}

export function ForgotPasswordForm({
  className,
  resetMode,
  onResetModeChange,
  ...props
}: ForgotPasswordFormProps) {
  const { t } = useTranslation()
  const [isLoading, setIsLoading] = useState(false)

  const { status } = useStatus()
  const phoneResetEnabled = !!status?.phone_verification

  const {
    isTurnstileEnabled,
    turnstileSiteKey,
    turnstileToken,
    setTurnstileToken,
    validateTurnstile,
  } = useTurnstile()
  const {
    secondsLeft,
    isActive,
    start: startCountdown,
  } = useCountdown({ initialSeconds: PASSWORD_RESET_COUNTDOWN })
  const { getCaptchaVerifyParam } = useAliyunCaptcha()
  const {
    isSending: isSendingSMS,
    secondsLeft: smsSecondsLeft,
    isActive: smsIsActive,
    sendCode: sendSMS,
  } = useSMSVerification({
    turnstileToken,
    validateTurnstile,
    getCaptchaVerifyParam,
  })

  const emailForm = useForm<z.infer<typeof forgotPasswordFormSchema>>({
    resolver: zodResolver(forgotPasswordFormSchema),
    defaultValues: { email: '' },
  })

  const phoneForm = useForm<z.infer<typeof phoneResetFormSchema>>({
    resolver: zodResolver(phoneResetFormSchema),
    defaultValues: { phone: '', code: '', password: '', confirmPassword: '' },
  })

  const turnstileReady = !isTurnstileEnabled || Boolean(turnstileToken)
  const phoneValue = phoneForm.watch('phone')

  async function onEmailSubmit(
    data: z.infer<typeof forgotPasswordFormSchema>
  ) {
    if (!validateTurnstile()) return

    setIsLoading(true)
    try {
      const res = await sendPasswordResetEmail(data.email, turnstileToken)
      if (res?.success) {
        emailForm.reset()
        startCountdown()
        toast.success(t('Reset email sent, please check your inbox'))
      } else {
        toast.error(res?.message || t('Failed to send reset email'))
      }
    } catch (_error) {
      // Errors are handled by global interceptor
    } finally {
      setIsLoading(false)
    }
  }

  async function onPhoneSubmit(data: z.infer<typeof phoneResetFormSchema>) {
    if (!validateTurnstile()) return

    setIsLoading(true)
    try {
      const res = await resetPasswordByPhone({
        phone: data.phone,
        code: data.code,
        password: data.password,
        turnstile: turnstileToken,
      })
      if (res?.success) {
        phoneForm.reset()
        toast.success(
          t('Password reset successful, please sign in with your new password')
        )
      } else {
        toast.error(res?.message || t('Failed to reset password'))
      }
    } catch (_error) {
      // Errors are handled by global interceptor
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div
      className={cn('grid gap-2', className)}
      {...(props as React.HTMLAttributes<HTMLDivElement>)}
    >
      {/* Reset mode toggle */}
      {phoneResetEnabled && (
        <div className='grid grid-cols-2 gap-1 rounded-lg bg-muted p-1'>
          <button
            type='button'
            onClick={() => onResetModeChange('email')}
            className={cn(
              'rounded-md px-3 py-1.5 text-sm font-medium transition-colors',
              resetMode === 'email'
                ? 'bg-background text-foreground shadow-sm'
                : 'text-muted-foreground hover:text-foreground'
            )}
          >
            {t('Reset via email')}
          </button>
          <button
            type='button'
            onClick={() => onResetModeChange('phone')}
            className={cn(
              'rounded-md px-3 py-1.5 text-sm font-medium transition-colors',
              resetMode === 'phone'
                ? 'bg-background text-foreground shadow-sm'
                : 'text-muted-foreground hover:text-foreground'
            )}
          >
            {t('Reset via SMS')}
          </button>
        </div>
      )}

      {resetMode === 'email' && (
        <Form {...emailForm}>
          <form
            onSubmit={emailForm.handleSubmit(onEmailSubmit)}
            className='grid gap-2'
          >
            <FormField
              control={emailForm.control}
              name='email'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder='name@example.com' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              type='submit'
              className='mt-2'
              disabled={isLoading || isActive || !turnstileReady}
            >
              {isActive
                ? t('Resend ({{seconds}}s)', { seconds: secondsLeft })
                : t('Send reset email')}
              {isLoading ? <Loader2 className='animate-spin' /> : <ArrowRight />}
            </Button>
          </form>
        </Form>
      )}

      {phoneResetEnabled && resetMode === 'phone' && (
        <Form {...phoneForm}>
          <form
            onSubmit={phoneForm.handleSubmit(onPhoneSubmit)}
            className='grid gap-2'
          >
            {/* Phone Field */}
            <FormField
              control={phoneForm.control}
              name='phone'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Phone number')}</FormLabel>
                  <FormControl>
                    <Input
                      placeholder={t('Enter phone number')}
                      type='tel'
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* SMS Code Field */}
            <FormField
              control={phoneForm.control}
              name='code'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('SMS verification code')}</FormLabel>
                  <div className='flex items-center gap-2'>
                    <FormControl>
                      <Input
                        placeholder={t('SMS verification code')}
                        {...field}
                      />
                    </FormControl>
                    <Button
                      variant='outline'
                      type='button'
                      disabled={
                        isLoading ||
                        isSendingSMS ||
                        smsIsActive ||
                        !phoneValue ||
                        !turnstileReady
                      }
                      onClick={() => sendSMS(phoneValue || '')}
                    >
                      {smsIsActive ? (
                        t('Resend ({{seconds}}s)', { seconds: smsSecondsLeft })
                      ) : isSendingSMS ? (
                        <Loader2 className='h-4 w-4 animate-spin' />
                      ) : (
                        t('Send SMS')
                      )}
                    </Button>
                  </div>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* New Password Field */}
            <FormField
              control={phoneForm.control}
              name='password'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('New password')}</FormLabel>
                  <FormControl>
                    <PasswordInput
                      placeholder={t('Enter password (8-20 characters)')}
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* Confirm Password Field */}
            <FormField
              control={phoneForm.control}
              name='confirmPassword'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{t('Confirm password')}</FormLabel>
                  <FormControl>
                    <PasswordInput
                      placeholder={t('Confirm password')}
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              type='submit'
              className='mt-2'
              disabled={isLoading || !turnstileReady}
            >
              {t('Reset password')}
              {isLoading ? <Loader2 className='animate-spin' /> : <ArrowRight />}
            </Button>
          </form>
        </Form>
      )}

      {isTurnstileEnabled && (
        <div className='mt-2'>
          <Turnstile siteKey={turnstileSiteKey} onVerify={setTurnstileToken} />
        </div>
      )}
    </div>
  )
}
