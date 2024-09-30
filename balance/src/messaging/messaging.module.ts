import { Module } from '@nestjs/common';
import { MessagingController } from './messaging.controller';
import { MessagingService } from './messaging.service';
import { DbModule } from '../db/db.module';

@Module({
  imports: [DbModule],
  controllers: [MessagingController],
  providers: [DbModule, MessagingService],
})
export class MessagingModule {}
