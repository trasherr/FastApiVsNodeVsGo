import { Sequelize } from "sequelize";

export const sequelize = new Sequelize("benchmark_db", "postgres", "postgres", {
    host: "localhost",
    port: 5432,
    dialect: "postgres",
    });

export default sequelize;