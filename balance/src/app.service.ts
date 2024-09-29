import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getBalance(): string {
    return 'Hello World!';
  }
}
