import { React, useRef, useState} from "react"


const Navbar = () => {

    const [state, setState] = useState(false)
    const navRef = useRef()

    // Replace # path with your path
    const navigation = [
        { 
            title: "About", 
            path: "/#About" 
        },
        { 
            title: "Our API", 
            path: "/swagger/index.html" 
        },
        /*{
            title : "Services",
            path : "/"
        },
        { 
            title: "Work with me", 
            path: "/" 
        },*/
    ]

    

    return (
        <nav ref={navRef} className="nav-ctn">
            <div className="nav-wrapper">
                <div className="">
                    <a href="/" className="text-2xl logo-name">
                        PowerNotes
                    </a>
                    <div className="lg:hidden">
                        <button className="outline-none p-2 rounded-md focus:border-gray-400 focus:border" onClick={() => setState(!state)}>
                            {
                                state ? (
                                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">
                                        <path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd" />
                                    </svg>
                                ) : (
                                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 8h16M4 16h16" />
                                    </svg>
                                )
                            }
                        </button>
                    </div>
                </div>
                <div className={`${ state ? 'h-screen pb-20 overflow-auto pr-4' : 'hidden'}`}>
                    
                    <div className="nav-side">
                        <ul className="">
                            {
                                navigation.map((item, idx) => {
                                    return (
                                        <li key={idx} className="nav-text">
                                            <a href={item.path}>
                                                { item.title }
                                            </a>
                                        </li>
                                    )
                                })
                            }
                            <li className="mt-8 lg:mt-0">
                                <a href="/signup" className="btn-primary  lg:inline">
                                    Signup
                                </a>
                                <a href="/login" className="btn-primary btn-primary--stand  lg:inline">
                                    Login
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </nav>
    )
}

export default Navbar;