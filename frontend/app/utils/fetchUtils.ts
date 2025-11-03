const url:string = process.env.BACKEND_URL ?? "http://localhost:8080/"
const register =()=>{
    fetch(url+"register")
}