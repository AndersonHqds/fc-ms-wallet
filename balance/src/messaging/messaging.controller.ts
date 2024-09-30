import { Controller } from '@nestjs/common';
import { EventPattern, MessagePattern, Payload, Transport } from '@nestjs/microservices';
import { MessagingService } from './messaging.service';

@Controller('messaging')
export class MessagingController {
  constructor(
    private readonly messagingService: MessagingService
  ) {}

  @EventPattern('balances', Transport.KAFKA)
  handleEvent(@Payload() payload: any): Promise<void> {
    console.log(`Received Message: ${JSON.stringify(payload, null, 2)}`)
    const data = payload.Payload;
    return this.messagingService.updateOrCreateBalance({
      from: {
        account_id: data.account_id_from,
        balance: data.balance_account_id_from
      },
      to: {
        account_id: data.account_id_to,
        balance: data.balance_account_id_to
      }
    })
  }
}
