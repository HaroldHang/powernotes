import { React } from "react";
import Navbar from "../components/navbar";
import Footer from "../components/footer";

const SignUp = () => {

    return (
        <div className="signup-page">
            <Navbar/>
            <main>

                <form className="formCtn">
                    <h1>Sign up here</h1>
                    <div className="formInput">
                        <input type="text" placeholder="Your Email"/>
                    </div>
                    <div className="formInput">
                        <input type="text" placeholder="Your Password"/>
                    </div>
                    <div className="submitInput">
                        <button className="btn-primary btn-primary--stand">
                            Sign Up
                        </button>
                    </div>
                </form>


            </main>
            <Footer/>
        </div>
    )

}

export default SignUp
