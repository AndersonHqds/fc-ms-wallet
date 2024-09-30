import { Inject, Injectable } from '@nestjs/common';

type GetBalanceByAccountReturnDTO = {
  balance: number;
}

type ErrorReturn = {
  error: string;
}

@Injectable()
export class AppService {
  constructor(
    @Inject('PG_CONNECTION')
    private readonly connection: any
  ) {}

  async getBalanceByAccountId(accountId: string): Promise<GetBalanceByAccountReturnDTO | ErrorReturn> {
    console.log({ accountId })
    const {rows} = await this.connection.query("SELECT * FROM balances WHERE account_id = $1", [accountId]);
    if (rows.length > 0) {
      return rows[0]
    }
    return {
      error: "This account not exists"
    };
  }
}
