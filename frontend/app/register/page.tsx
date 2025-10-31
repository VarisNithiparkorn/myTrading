"use client"
import Header from "../Components/Header"
import Image from "next/image"
import HomeEgg from "../Components/images/egg_home.png"
import { useState } from "react"
export default function Register() {
    const [emailFocus,setEmailFocusStatus]= useState<boolean>(false)
    const [pwdFocus,setPwdFocus] = useState<boolean>(false)
    const setDropDown=(section:string) =>{
    const services : string[] = ["ตลาดซื้อ-ขาย","ซื้อ-ขาย คริปโตเคอร์เรนซี","ซื้อ Bitcoin","ซื้อ Ethereum","ซื้อ Solana"]
    const aboutUs : string[] = ["รู้จักเรา","ภารกิจของเรา", "ข้อตกลงผู้ใช้บริการ","คำสั่งซื้อขายที่ไม่เหมาะสม","นโยบายการแจ้งเบาะแสและข้อร้องเรียน"]
    const help : string[] = ["คำถามที่พบบ่อย","ติดต่อเจ้าหน้าที่","ถวามปลอดภัย"]
    const annoucments : string[] = ["บล็อก","ประกาศ","สื่อมวลชน","BitKai API"]
    const others : string[]=["ติดต่อเรา","ร่วมงานกับเรา","ค่าธรรมเนียม","ติดต่อด้านพัฒนาธุรกิจ"]
    const dropdownLayout:string = "mt-2 flex flex-col"
    const textDropdown:string =" text-gray-400 "
        if(section === "services"){
            return <div className={dropdownLayout}>{services.map((s)=> <a href="" className={textDropdown}>{s}</a>)}</div>
        }else if(section === "aboutUs"){
            return <div className={dropdownLayout}>{aboutUs.map((a)=><a className={textDropdown}>{a}</a>)}</div>
        }else if(section === "help"){
            return <div className={dropdownLayout}>{help.map((h)=> <a className={textDropdown}>{h}</a>)}</div> 
        }else if(section === "annoucement"){
            return <div className={dropdownLayout}>{annoucments.map((a)=> <a className={textDropdown}>{a}</a>)}</div> 
        }else if(section === "others"){
            return <div className={dropdownLayout}>{others.map((o)=><a className={textDropdown}>{o}</a>)}</div>
        }
    }
    const setTypingStatus =(e: React.FocusEvent<HTMLInputElement>)=>{
        const id = e.target.id
        if(id === "email"){
            setEmailFocusStatus(!emailFocus)
        }else if(id === "password"){
            setPwdFocus(!pwdFocus)
        }
    }
    return(
        <div >
            <Header>
                
            </Header>
            <div className="  min-h-screen bg-amber-50">
                <div className=" flex flex-wrap w-[90%] h-[600px] mx-auto">
                    <div className="  max-[1025px]:hidden mt-20 ml-20 flex-1 flex flex-col">
                        <p className=" text-amber-400 mb-2">
                            ยินดีต้อนรับสู่บิทKai เอ็กซ์เชนจ์
                        </p>
                        <div className=" w-10 h-2 mb-2 bg-amber-400 rounded-2xl">

                        </div>
                        <p className=" text-gray-400  font-thin w-1/2">
                            เพิ่มโอกาสในการลงทุนกับคริปโตมากกว่า 3 เหรียญ บนแพลตฟอร์มที่คุณซื้อขายได้ด้วยเงินบาท
                        </p>
                        <Image className=" w-[400px] h-[400px]" src={HomeEgg}
                            alt="HomeEgg">

                        </Image>
                    </div>
                    <div className=" mt-20 flex flex-col  flex-1">
                        <h1 
                        className="font-custom-font max-[1025px]:w-full text-3xl w-[80%] text-center mb-5">
                            เข้าสู่ระบบ
                        </h1>
                        <div className="max-[1025px]:w-full w-[80%]  relative">
                             <p className=" font-light text-gray-500 text-[14px]">
                                อีเมล
                            </p>
                            <input id="email" onBlur={setTypingStatus} onFocus={setTypingStatus} type="email" placeholder={emailFocus ? "" : "กรอกอีเมล"} className=" p-5 pl-9 text-[14px] rounded focus:border-amber-400 focus:outline-none focus:border-2  border-gray-400 bg-white border w-full h-10">
                            
                            </input>
                            <img src={"/mailIcon/"+ (emailFocus ? "mail_yellow.png":"mail.png")} className=" opacity-50 w-[15px] h-[15px] absolute bottom-[12px] left-3"></img>
                        </div>
                        <div className="max-[1025px]:w-full w-[80%] mt-2  relative">
                             <p className=" font-light text-gray-500 text-[14px]">
                                รหัสผ่าน
                            </p>
                            <input onBlur={setTypingStatus} onFocus={setTypingStatus} id="password" type="password" placeholder="กรอกรหัสผ่าน" className=" p-5 pl-9 text-[14px] rounded focus:border-amber-400 focus:outline-none focus:border-2  border-gray-400 bg-white border w-full h-10">
                            
                            </input>
                            <img src={"/" + (pwdFocus ? "lock_yellow.png":"padlock.png") }className=" opacity-50 w-[15px] h-[15px] absolute bottom-[12px] left-3"></img>
                        </div>
                        <p className="max-[1025px]:w-full text-amber-400 w-[80%] text-right">
                            ลืมรหัสผ่าน?
                        </p>
                        <button className="max-[1025px]:w-full font-extrabold  mt-5 bg-amber-400 text-white rounded border w-[80%] h-12" >
                            เข้าสู่ระบบ
                        </button>
                        <div className="max-[1025px]:w-full mt-6 mb-6 relative w-[80%] border border-amber-400">
                            <p className="  bg-amber-50 absolute top-1/2 left-1/2 transform text-amber-400 -translate-x-1/2 -translate-y-1/2">
                                หรือ
                            </p>
                        </div> 
                        <button className="max-[1025px]:w-full text-amber-400 border w-[80%] h-12 rounded border-amber-400" >
                            เข้าสู่ระบบ
                        </button>   
                    </div>
                </div>
                <div className=" max-[1025px]:flex-col mx-auto flex flex-5 w-[80%] flex-row">
                    <div className=" flex-1">
                        <h4>
                            บริการของเรา
                        </h4>
                        {setDropDown("services")}
                    </div>
                    <div className=" flex-1">
                        <h4>
                            เกี่ยวกับbitKai
                        </h4>
                        {setDropDown("aboutUs")}
                    </div>
                    <div className=" flex-1">
                        <h4>
                            ช่วยเหลือ
                        </h4>
                        {setDropDown("help")}
                    </div>
                    <div className=" flex-1">
                        <h4>
                            ประกาศและข่าวสาร
                        </h4>
                        {setDropDown("annoucement")}
                    </div>
                    <div className=" flex-1">
                        <h4>
                            อื่นๆ
                        </h4>
                        {setDropDown("others")}
                    </div>
                </div>
            </div>
        </div>
    )
    
}