import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MessagingModule } from './messaging/messaging.module';
import { DbModule } from './db/db.module';

@Module({
  imports: [MessagingModule, DbModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
