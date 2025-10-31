import Image from "next/image"
import Logo from "./images/Bitkai_logo.png"
import BurgerBar from "./images/burger-bar.png"
export default function Header(){
    return(
        <div className="  h-[60px] bg-white" >
            <div className=" p-2 w-[80%]  h-full mx-auto flex justify-between items-center">
                <div className=" w-[120px] ">
                    <Image className=" w-full h-[80%]" src={Logo}
                    alt="logo">

                    </Image>
                </div>
                <div className=" h-full  flex items-center">
                    <button className=" h-[30px] w-[100px] mr-2" >
                        <span className=" text-amber-400 text-[16px]">
                            เข้าสู่ระบบ
                        </span>
                    </button>
                    <button className=" bg-amber-400 h-[30px] w-[120px] rounded">
                        <span className=" text-[16px] text-white">
                            สมัครใช้งานฟรี
                        </span>
                    </button>
                </div>
            </div>
        </div>
    )
}