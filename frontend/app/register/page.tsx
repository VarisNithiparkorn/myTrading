"use client"
import Header from "../Components/Header"
import Image from "next/image"
import HomeEgg from "../Components/images/egg_home.png"
import Caution from "../Components/images/caution.png"
import RegisterEgg from "../Components/images/register_egg.png"
import { useState } from "react"
export default function Register() {
    const [email,setEmail]= useState<string>('')
    const [password,setPassword] = useState<string>('')
    const [confirmPassword, setConfirmPassword]= useState<string>('')
    const [emailFocus,setEmailFocusStatus]= useState<boolean>(false)
    const [pwdFocus,setPwdFocus] = useState<boolean>(false)
    const [confirmPwdFocus,setConfirmPwdFocus] = useState<boolean>(false)
    const [isValidLength,setIsValidLength] = useState<boolean>(false)
    const [isValidCharacter,setIsValidCharacter] = useState<boolean>(false)
    const [isValidNumber,setIsValidNumber] = useState<boolean>(false)
    const [isValidSymbol,setIsValidSymbol] = useState<boolean>(false)
    const [showPwdErr,setShowPwdErr] = useState<boolean>(false)
    const [isClicked,setIsClicked] = useState<boolean>(false)
    const [pwdErrMsg,setPwdErrMsg] = useState<string>("")
    const [currentForm , setCurrentForm] = useState<string>("signIn")
    const setDropDown=(section:string) =>{
    const services : string[] = ["ตลาดซื้อ-ขาย","ซื้อ-ขาย คริปโตเคอร์เรนซี","ซื้อ Bitcoin","ซื้อ Ethereum","ซื้อ Solana"]
    const aboutUs : string[] = ["รู้จักเรา","ภารกิจของเรา", "ข้อตกลงผู้ใช้บริการ","คำสั่งซื้อขายที่ไม่เหมาะสม","นโยบายการแจ้งเบาะแสและข้อร้องเรียน"]
    const help : string[] = ["คำถามที่พบบ่อย","ติดต่อเจ้าหน้าที่","ถวามปลอดภัย"]
    const annoucments : string[] = ["บล็อก","ประกาศ","สื่อมวลชน","BitKai API"]
    const others : string[]=["ติดต่อเรา","ร่วมงานกับเรา","ค่าธรรมเนียม","ติดต่อด้านพัฒนาธุรกิจ"]
    const dropdownLayout:string = "mt-2 flex flex-col"
    const textDropdown:string =" text-gray-600 hover:text-gray-800"
        if(section === "services"){
            return <div className={dropdownLayout}>{services.map((s,i)=> <a key={i} href="" className={textDropdown}>{s}</a>)}</div>
        }else if(section === "aboutUs"){
            return <div className={dropdownLayout}>{aboutUs.map((a,i)=><a key={i} href="" className={textDropdown}>{a}</a>)}</div>
        }else if(section === "help"){
            return <div className={dropdownLayout}>{help.map((h,i)=> <a key={i} href="" className={textDropdown}>{h}</a>)}</div> 
        }else if(section === "annoucement"){
            return <div className={dropdownLayout}>{annoucments.map((a,i)=> <a key={i} href="" className={textDropdown}>{a}</a>)}</div> 
        }else if(section === "others"){
            return <div className={dropdownLayout}>{others.map((o,i)=><a key={i} href="" className={textDropdown}>{o}</a>)}</div>
        }
    }
    const clearForm = ():void=>{
        setIsClicked(false)
        setEmail('')
        setPassword('')
        setConfirmPassword('')
    }
    const setTypingStatus =(e: React.FocusEvent<HTMLInputElement>):void=>{
        const id = e.target.id
        if(id === "email"){
            setEmailFocusStatus(!emailFocus)
        }else if(id === "password"){
            setPwdFocus(!pwdFocus)
        }else if(id === "confirmPassword"){
            setConfirmPwdFocus(!confirmPwdFocus)
        }
    }
    const handleFormChange = (formType:"signIn"|"register"):void=>{
        setCurrentForm(formType)
        setIsClicked(false)
        clearForm()
    }
    const setClick=()=>{
        setIsClicked(true)
    }
    const validateInput=(inputType:"email"|"password"|"confirmPassword",input:string)=>{
        if(inputType ==="email"){

        }else if(inputType === "password"){
            if(input.length === 0){
                setShowPwdErr(false)
                return
            }
            const passwordCharacterCase = new RegExp(
            `(?=.*[a-z])` + 
            `(?=.*[A-Z])` 
            );
            const passwordNumberCase = new RegExp( `(?=.*[0-9])` )
            const passwordSymbolCase = new RegExp(`(?=.*[!@#$%^&*\\-_+=])`)
            setIsValidLength(input.length+1 >= 12)
            setIsValidCharacter(passwordCharacterCase.test(input))
            setIsValidNumber(passwordNumberCase.test(input))
            setIsValidSymbol(passwordSymbolCase.test(input))
            if(isValidLength && isValidCharacter&& isValidNumber&& isValidSymbol){
                setPwdErrMsg("")
                setShowPwdErr(false)
            }else{
                setPwdErrMsg("โปรดตั้งรหัสผ่านตามเงื่อนไขด้านล่างเพื่อความปลอดภัย")
                setShowPwdErr(true)
            }
        }else if(inputType === "confirmPassword"){

        }
    }
    const intitialInput=(e:React.FocusEvent<HTMLInputElement>)=>{
        const value: string = e.target.value
        const id: string = e.target.id
        if(id === "email"){
            setEmail(value)
        }else if(id === "password"){
            validateInput(id,value)
            setPassword(value)
        }else if(id ==="confirmPassword"){
            setConfirmPassword(value)
        }
    }
    const sendAuthentication = ():void=>{
        if(currentForm ==="register"){
            
        }else if(currentForm === "signIn"){

        }
    }
    return(
        <div >
            <Header onFormChange={handleFormChange}>
                
            </Header>
            <div className="  min-h-screen bg-amber-50">
                <div className=" flex flex-wrap w-[90%] h-[600px] mx-auto">
                    <div className="  max-[1025px]:hidden mt-20 ml-20 flex-1 flex flex-col">
                        <p className=" text-amber-400 mb-2">
                            { currentForm === "signIn"?"ยินดีต้อนรับสู่บิทKai เอ็กซ์เชนจ์":"สมัครสมาชิก"}
                        </p>
                        <div className=" w-10 h-2 mb-2 bg-amber-400 rounded-2xl">

                        </div>
                        <p className=" text-gray-600  font-thin w-1/2">
                            {currentForm === "signIn"?"เพิ่มโอกาสในการลงทุนกับคริปโตมากกว่า 3 เหรียญ บนแพลตฟอร์มที่คุณซื้อขายได้ด้วยเงินบาท":"แพลตฟอร์มซื้อขายคริปโตที่มีผู้สมัครใช้งานมากกว่า 1 บัญชี และมีมูลค่าการซื้อขายสูงเป็นอันดับ 2 ของไทย"}
                        </p>
                        <Image className=" w-[400px] h-[400px]" src={currentForm ==="signIn"?HomeEgg:RegisterEgg}
                            alt="HomeEgg">

                        </Image>
                    </div>
                    <div className=" mt-20 flex flex-col  flex-1">
                        <h1 
                        className="font-custom-font max-[1025px]:w-full text-3xl w-[80%] text-center mb-5">
                            {currentForm ==="signIn"?"เข้าสู่ระบบ":"สมัครสมาชิก"}
                        </h1>
                        <div className="max-[1025px]:w-full w-[80%]  relative">
                             <p className=" font-light text-gray-500 text-[14px]">
                                อีเมล
                            </p>
                            <input id="email" value={email} onInput={intitialInput} onBlur={setTypingStatus} onFocus={setTypingStatus} type="email" placeholder={emailFocus ? "" : "กรอกอีเมล"} className=" p-5 pl-9 text-[14px] rounded hover:border-amber-400 focus:border-amber-400 focus:outline-none focus:border-2  border-gray-400 bg-white border w-full h-10">
                            
                            </input>
                            <img src={"/mailIcon/"+ (emailFocus  ? "mail_yellow.png":"mail.png")} className=" opacity-50 w-[15px] h-[15px] absolute bottom-[12px] left-3"></img>
                        </div>
                        <div className="max-[1025px]:w-full w-[80%] mt-2  relative">
                             <p className=" font-light text-gray-500 text-[14px]">
                                รหัสผ่าน
                            </p>
                            <input onClick={setClick} value={password} onInput={intitialInput}  onBlur={setTypingStatus} onFocus={setTypingStatus} id="password" type="password" placeholder={pwdFocus?"":"กรอกรหัสผ่าน"} className={`${showPwdErr ? "border-red-600" : "border-gray-400"} 
                            hover:border-amber-400 
                             p-5 pl-9 text-[14px] rounded 
                            focus:border-amber-400 focus:outline-none focus:border-2
                            bg-white border w-full h-10
                            `}>
                            
                            </input>
                            <img src={"/lock/" + (pwdFocus ? "lock_yellow.png": !pwdFocus && showPwdErr ? "padlock_red.png":"padlock.png") }className=" opacity-50 w-[15px] h-[15px] absolute bottom-[12px] left-3"></img>
                        </div>
                        <p className="m-1 text-xs">
                            {password.length !== 0 ? pwdErrMsg:""}
                        </p>
                        { isClicked && currentForm === "register" ?
                        <ul className=" p-3 bg-amber-100 w-[80%] mt-2">
                            <li className=" m-1 flex flex-row items-center">
                                <img src={"/checklist/"+(isValidLength && password.length !== 0 ?"check_yellow.png": password.length === 0 ?"check.png":"incorrect.png")} className=" w-4 h-4"></img>
                                <span className=" ml-3">มีความยาวอย่างน้อย 12 ตัว</span>
                            </li>
                            <li className="m-1 flex flex-row items-center">
                                <img src={"/checklist/"+(isValidCharacter && password.length !== 0 ?"check_yellow.png":password.length === 0 ?"check.png":"incorrect.png")} className=" w-4 h-4"></img>
                                <span className=" ml-3">มีตัวอักษรภาษาอังกฤษพิมพ์ใหญ่และพิมพ์เล็ก</span>
                            </li>
                            <li className="m-1 flex flex-row items-center">
                                <img src={"/checklist/"+(isValidNumber && password.length !== 0 ?"check_yellow.png":password.length === 0 ?"check.png":"incorrect.png")} className=" w-4 h-4"></img>
                                <span className=" ml-3">มีตัวเลข 0-9 อย่างน้อย 1 ตัว</span>
                            </li>
                            <li className="m-1 flex flex-row items-center">
                                <img src={"/checklist/"+(isValidSymbol && password.length !== 0 ?"check_yellow.png":password.length === 0 ?"check.png":"incorrect.png")} className=" w-4 h-4"></img>
                                <span className=" ml-3">มีสัญลักษณ์ !@#$%^&*-_+= อย่างน้อย 1 ตัว</span>
                            </li>
                        </ul>
                        : null}
                        {currentForm ==="signIn"?
                        null
                        :<div className="max-[1025px]:w-full w-[80%] mt-2  relative">
                             <p className=" font-light text-gray-500 text-[14px]">
                                ยืนยันรหัสผ่าน
                            </p>
                            <input  onBlur={setTypingStatus} value={confirmPassword} onInput={intitialInput} onFocus={setTypingStatus} id="confirmPassword" type="password" placeholder={confirmPwdFocus ? " ":"ยืนยันรหัสผ่าน"} className="hover:border-amber-400 p-5 pl-9 text-[14px] rounded focus:border-amber-400 focus:outline-none focus:border-2  border-gray-400 bg-white border w-full h-10">
                            
                            </input>
                            <img src={"/lock/" + (confirmPwdFocus ? "lock_yellow.png":"padlock.png") }className=" opacity-50 w-[15px] h-[15px] absolute bottom-[12px] left-3"></img>
                        </div>}
                        {currentForm === "signIn"?
                        <p className="max-[1025px]:w-full hover:cursor-pointer text-amber-400 w-[80%] text-right">
                            ลืมรหัสผ่าน?
                        </p>:null
                        }
                       
                        <button onClick={sendAuthentication} className="max-[1025px]:w-full font-extrabold hover:cursor-pointer  mt-5 bg-amber-400 hover:bg-amber-500 hover: text-white rounded border w-[80%] h-12" >
                            {currentForm === "signIn"?"เข้าสู่ระบบ":"สมัครสมาชิก"}
                        </button>
                        {
                            currentForm === "signIn" && (
                        <div>
                        <div className="max-[1025px]:w-full mt-6 mb-6 relative w-[80%] border border-gray-400">
                            <p className="  bg-amber-50 absolute top-1/2 left-1/2 transform text-gray-400 -translate-x-1/2 -translate-y-1/2">
                                หรือ
                            </p>
                        </div> 
                        <button className=" hover:cursor-pointer max-[1025px]:w-full text-amber-400 border w-[80%] h-12 rounded border-amber-400" >
                            เข้าสู่ระบบ ด้วย Passkey
                        </button>
                        </div>)
                        }   
                    </div>
                </div>
                <div className=" mt-24 max-[1025px]:flex-col mx-auto flex flex-5 w-[80%] flex-row">
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