import { User } from "./models.js";

export const Login = async (req, res) => {
    let user = await User.findOne({ where: { email: req.body.email, password: req.body.password } })
     if(user){
        res.send(user);
    } else {
        res.status(404).send({"message": "User not found"});
    }
};

export const Create = async (req, res) => {
    User.create({
        name: req.body.name,
        email: req.body.email,
        password: req.body.password,
        age: req.body.age
    }).then(user => {
        res.send(user);
    }).catch(err => {
        console.log(err);
        
        res.status(500).send({"message": "User exists"});
    });
};

export const Update = async (req, res) => {
    User.update({
        name: req.body.name,
        password: req.body.password,
        age: req.body.age
    }, {
        where: {
            email: req.body.email
        }
    }).then(user => {
        if(user[0] === 0) {
            res.status(404).send({"message": "User not found"});
            return ;
        }
        res.send(user);
    })
};

export const Delete = async (req, res) => { 
    let user = await User.findOne({ where: {  email: req.body.email } });
    if( user ){
        await User.destroy({
            where: {
                email: req.body.email
            }
        })
        res.send(user);
        return ;
    }
    res.status(404).send({"message": "User not found"});
};

export const GetAll = async (req, res) => {
    let users = await User.findAll();
    res.send(users);
};