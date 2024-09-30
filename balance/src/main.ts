import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { Transport } from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.connectMicroservice({
    name: 'KAFKA_SERVICE',
    transport: Transport.KAFKA,
    options: {
      client: {
        clientId: 'wallet-balance',
        brokers: ['kafka:29092']
      },
      consumer: {
        groupId: 'wallet'
      }
    }
  })
  await app.startAllMicroservices();
  await app.listen(3003);
}
bootstrap();
