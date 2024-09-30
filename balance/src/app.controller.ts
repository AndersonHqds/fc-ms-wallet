import { Controller, Get, HttpStatus, Param, Res } from '@nestjs/common';
import { AppService } from './app.service';
import { Response } from 'express';

@Controller('balances')
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get(':account_id')
  async getBalance(@Param('account_id') id: string, @Res() res: Response) {
    const balance = await this.appService.getBalanceByAccountId(id);
    return res.status(HttpStatus.OK).json(balance)
  }
}
