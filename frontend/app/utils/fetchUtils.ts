import { json } from "stream/consumers"
import { Account } from "./interface"

const url:string = process.env.BACKEND_URL ?? "http://localhost:5050/"
export const register = async(account:Account)=>{
    try{
    const res:Response = await fetch(url+"api/v1/register",{
        method:"POST",
        body:JSON.stringify(account)
    })
    return res.status
    }catch(e){
        console.log("register faile")
    }
}