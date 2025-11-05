"use client"
import { useRouter, useSearchParams } from "next/navigation";
import Header from "../Components/Header";
import Image from "next/image";
import verifyLogo from "../Components/images/egg_verifymail.png"
import { useEffect, useState } from "react";
import { verifyEmail } from "../utils/fetchUtils";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
export default function VerifyEmail(){
    const router:AppRouterInstance = useRouter()
    const searchParam = useSearchParams()
    const [status, setStatus] = useState<number>(0)
    const handleFormChange = (formType:"signIn"|"register"):void=>{
        router.push("/register")
    }
    useEffect(()=>{
        const token:string|null =  searchParam.get("token")
        const sendToken = async()=>{
            if(token === null){
            return
        }
            const status = await verifyEmail(token)
            setStatus(status)
        }
        sendToken()
    },[])


    return(
        <div>
            <Header onFormChange={handleFormChange}>
                            
            </Header>
            {
                status === 200 ?
                <div className=" flex justify-center items-center  min-h-screen w-full bg-amber-50">
                    <div className=" h-full flex flex-col justify-center items-center min-w-[50%] w-full:">
                        <Image className=" max-w-[250px] max-h-[250px]" src={verifyLogo} alt="verify success">
                        </Image>
                        <h2 className=" max-[426px]:text-[12px] max-[769px]:text-[18px] m-3 text-2xl ">
                            Your signup has been confirmed successfully via email
                        </h2>
                        <p className="m-1 text-[14px] max-[769px]:text-[12px] max-[426px]:text-[8px]">
                            you can now access Bitkai Exchange     
                        </p>
                        <button className=" hover:bg-amber-500 m-6 w-3/8 h-[35px] rounded bg-amber-400">
                            <span className=" text-white"> 
                                OK
                            </span>
                        </button>
                    </div>
                </div>
                : null
            }
        </div>
    )
}