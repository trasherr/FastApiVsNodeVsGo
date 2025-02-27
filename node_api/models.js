import sequelize from "./db.js";
import { DataTypes } from "sequelize";

export const User = sequelize.define('users', {
    id: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    name: DataTypes.STRING,
    email: { type: DataTypes.STRING, unique: true},
    password: DataTypes.STRING,
    age: DataTypes.INTEGER
  },
  {
    timestamps: false
  }
);

// sequelize.sync({ force: true }).then(() => {
//     console.log('Database & tables created!');
//     for(let i = 0; i < 10_000; i++){
//       User.create({
//         name: `test${i}`,
//         email: `test${i}@gmail.com`,
//         password: `password${i}`,
//         age: Math.floor(Math.random() * 100)
//     })
//   }
  
// }); 


