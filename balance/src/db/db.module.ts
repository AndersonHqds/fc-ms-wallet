import { Module } from '@nestjs/common';
import { Pool } from "pg";

const DbProvider = {
  provide: 'PG_CONNECTION',
  useValue: new Pool({
    user: "postgres",
    host: "postgres",
    database: "wallet-balance",
    password: "root",
    port: 5432,
  })
}


@Module({
  providers: [DbProvider],
  exports: [DbProvider]
})
export class DbModule {}
