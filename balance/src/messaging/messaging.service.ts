import { Inject, Injectable } from '@nestjs/common';

type AccountBalance = {
  account_id: string;
  balance: number;
}

type AccountBalanceUpdate = {
  from: AccountBalance;
  to: AccountBalance;
}

@Injectable()
export class MessagingService {
  constructor(
    @Inject('PG_CONNECTION')
    private readonly connection: any
  ) {}

  async updateOrCreateBalance({from, to}: AccountBalanceUpdate) {
    try {
      console.log({ from, to })
      await this.connection.query(`BEGIN;`);
      await this.connection.query(`
          INSERT INTO balances (account_id, balance) VALUES ($1, $2)
          ON CONFLICT (account_id)
          DO UPDATE SET
            balance = EXCLUDED.balance;
      `, [from.account_id, from.balance ]);
      await this.connection.query(
        `
          INSERT INTO balances (account_id, balance) VALUES ($1, $2)
          ON CONFLICT (account_id)
          DO UPDATE SET
            balance = EXCLUDED.balance;
        `, [to.account_id, to.balance]
      )
      await this.connection.query(`COMMIT;`)
    }
    catch(e) {
      console.log(e)
      await this.connection.query('ROLLBACK');
    }
  }
}
