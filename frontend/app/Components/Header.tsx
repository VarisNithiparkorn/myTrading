import Image from "next/image"
import Logo from "./images/Bitkai_logo.png"
 interface HeaderProps {
    onFormChange: (formType: 'signIn' | 'register') => void;
}
export default function Header({onFormChange} : HeaderProps){

    return(
        <div className="  h-[60px] bg-white" >
            <div className=" p-2 w-[80%]  h-full mx-auto flex justify-between items-center">
                <div className=" w-[120px] ">
                    <Image className=" w-full h-[80%]" src={Logo}
                    alt="logo">

                    </Image>
                </div>
                <div className=" h-full  flex items-center">
                    <button onClick={()=> onFormChange('signIn')} className="hover:cursor-pointer h-[30px] w-[100px] mr-2" >
                        <span className=" text-amber-400 text-[16px] hover:text-amber-300">
                            เข้าสู่ระบบ
                        </span>
                    </button>
                    <button onClick={()=>onFormChange('register')} className="hover:cursor-pointer hover:bg-amber-500 bg-amber-400 h-[30px] w-[120px] rounded">
                        <span className=" text-[16px] text-white">
                            สมัครใช้งานฟรี
                        </span>
                    </button>
                </div>
            </div>
        </div>
    )
}