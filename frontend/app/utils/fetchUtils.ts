import { Account } from "./interface"

const url:string = process.env.BACKEND_URL ?? "http://localhost:5050/"
export const register = async(account:Account):Promise<number>=>{
    try{
    const res:Response = await fetch(url+"api/v1/register",{
        method:"POST",
        body:JSON.stringify(account)
    })
    return res.status
    }catch(e){
        console.log("register fail")
    }
    return 999 
}
export const login = async(account:Account):Promise<object|number>=>{
    try{
    const res:Response = await fetch(url+"api/v1/login",{
        method:"POST",
        body:JSON.stringify(account)
    })
    if (res.status !== 200) {
        return res.status
    }
    const data = res.json()
    return data
    }catch(e){
        console.log("register fail")
    } 
    return 999
}
export const verifyEmail = async(token:string):Promise<number>=>{
    try{
    const res:Response = await fetch(url+"api/v1/register?token="+token,{
        method:"GET"
    })
    return res.status
    }catch(e){
        console.log("verifymail fail")
    }
    return 999
}