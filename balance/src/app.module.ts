import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  imports: [ClientsModule.register([
    {
      name: 'KAFKA_SERVICE',
      transport: Transport.KAFKA,
      options: {
        client: {
          clientId: 'wallet-balance',
          brokers: ['localhost:9092']
        },
        consumer: {
          groupId: 'wallet'
        }
      }
    }
  ])],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
