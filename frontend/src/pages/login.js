import { React } from "react";
import Navbar from "../components/navbar";
import Footer from "../components/footer";

const Login = () => {

    return (
        <div className="login-page">
            <Navbar/>
            <main>

                <form className="formCtn">
                    <h1>Happy to see you. Sign in</h1>
                    <div className="formInput">
                        <input type="text" placeholder="Email"/>
                    </div>
                    <div className="formInput">
                        <input type="text" placeholder="Password"/>
                    </div>
                    <div className="submitInput">
                        <button className="btn-primary btn-primary--stand">
                            Sign in
                        </button>
                        
                    </div>
                </form>


            </main>
            <Footer/>
        </div>
    )

}

export default Login
