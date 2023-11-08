const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());

app.get("/",(req,res)=>{
    res.json([
        {
            name:"Sumit sapkota",
            age:25
        },
        {
            name:"Ram Prasad",
            age:26
        }
    ])
})

app.listen(4000,()=>{
    console.log("Listening on port 4000")
})