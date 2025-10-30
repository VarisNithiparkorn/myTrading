import Image from "next/image"
import Logo from "./images/Bitkai_logo.png"
export default function Header(){
    return(
        <div className=" w-screen h-[60px] bg-white" >
            <div className=" w-[80%] border h-full mx-auto flex">
                <div className=" w-[150px] h-full ">
                    <Image className=" w-full h-full" src={Logo}
                    alt="logo">

                    </Image>
                </div>
            </div>
        </div>
    )
}