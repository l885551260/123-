import { QRCodeSVG } from 'qrcode.react'
import { Loader2, X } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'

interface QrcodePaymentDialogProps {
  open: boolean
  qrCode: string | null
  amount?: number
  paymentLabel?: string
  polling?: boolean
  onClose: () => void
}

/**
 * QR Code Payment Dialog
 * Displays a QR code for Alipay/WeChat Pay native payment.
 * Polls for payment completion in the background.
 */
export function QrcodePaymentDialog({
  open,
  qrCode,
  amount,
  paymentLabel = '扫码支付',
  polling = false,
  onClose,
}: QrcodePaymentDialogProps) {
  const { t } = useTranslation()

  return (
    <Dialog open={open} onOpenChange={(v) => !v && onClose()}>
      <DialogContent className="sm:max-w-[360px]">
        <DialogHeader>
          <DialogTitle className="text-center">{paymentLabel}</DialogTitle>
        </DialogHeader>

        <div className="flex flex-col items-center gap-4 py-4">
          {amount !== undefined && amount > 0 && (
            <div className="text-2xl font-bold">
              ¥{amount.toFixed(2)}
            </div>
          )}

          {qrCode ? (
            <div className="rounded-lg border p-4 bg-white">
              <QRCodeSVG
                value={qrCode}
                size={200}
                level="M"
                includeMargin={false}
              />
            </div>
          ) : (
            <div className="flex h-[200px] w-[200px] items-center justify-center rounded-lg border">
              <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
            </div>
          )}

          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            {polling && <Loader2 className="h-4 w-4 animate-spin" />}
            <span>
              {polling
                ? t('Waiting for payment confirmation...')
                : t('Please scan the QR code to pay')}
            </span>
          </div>

          <p className="text-xs text-muted-foreground text-center">
            {t('The page will automatically refresh after payment is confirmed')}
          </p>
        </div>

        <button
          onClick={onClose}
          className="absolute right-4 top-4 rounded-sm opacity-70 transition-opacity hover:opacity-100"
        >
          <X className="h-4 w-4" />
        </button>
      </DialogContent>
    </Dialog>
  )
}
