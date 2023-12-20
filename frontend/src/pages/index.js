import { React } from "react";
import Navbar from "../components/navbar";
import Footer from "../components/footer";
import bannerSvg from "../assets/img/banner-svg.svg"
import integrationSvg from '../assets/img/integration.svg'
import notesSvg from '../assets/img/notes.svg'
import analyticsSvg from '../assets/img/analytics.svg'
const Home = () => {

    return (
        <div className="home-page">
            <Navbar/>
            <main className="">
                <section className="banner-hero">
                    <h1>
                        A Note Management <span className="note"> API<img src={bannerSvg} alt=""/></span>  easy to integate
                    </h1>
                        
                    <p>
                        Integrate our best api note management for all purpose : elearning platforms, blogs, training programs and more
                    </p>
                    <div className="btn-actions">
                    <a href="/" className="btn-primary btn-primary--stand  lg:inline">Get Started</a>
                    <a href="/" className="btn-primary">Read our API</a>
                    </div>

                </section>

                <section className="features">
                    <h3>Benefits for using our API</h3>

                    <div className="features-ctn">
                        <div className="aFeature">
                            <div className="imgCtn">
                                <img src={integrationSvg} alt="integration"/>
                            </div>
                            <h3>Multi Platforms</h3>
                            <p>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime mollitia, molestiae quas vel sint commodi repudiandae consequuntur voluptatum laborum numquam blanditiis harum quisquam eius sed odit fugiat iusto fuga praesentium optio, eaque rerum! Provident similique accusantium nemo autem. Veritatis obcaecati tenetur iure eius earum ut molestias architecto voluptate aliquam nihil
                            </p>
                        </div>
                        <div className="aFeature">
                            <div className="imgCtn">
                                <img src={notesSvg} alt="notes"/>
                            </div>
                            <h3>Unlimited Notes</h3>
                            <p>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime mollitia, molestiae quas vel sint commodi repudiandae consequuntur voluptatum laborum numquam blanditiis harum quisquam eius sed odit fugiat iusto fuga praesentium optio, eaque rerum! Provident similique accusantium nemo autem. Veritatis obcaecati tenetur iure eius earum ut molestias architecto voluptate aliquam nihil
                            </p>
                        </div>
                        <div className="aFeature">
                            <div className="imgCtn">
                                <img src={analyticsSvg} alt="analytics"/>
                            </div>
                            <h3>In Built Tracking</h3>
                            <p>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime mollitia, molestiae quas vel sint commodi repudiandae consequuntur voluptatum laborum numquam blanditiis harum quisquam eius sed odit fugiat iusto fuga praesentium optio, eaque rerum! Provident similique accusantium nemo autem. Veritatis obcaecati tenetur iure eius earum ut molestias architecto voluptate aliquam nihil
                            </p>
                        </div>
                    </div>
                </section>
                <section className="aboutUs" id="About">
                    <h3>About</h3>

                    <div className="aboutCtn">
                        <p>
                            The note management API is a great integration tools for managing notes in any system. This project is still actually under development and I'm welcoming anyone willing to contribute to expand many functionnalities.  If you're interest visit my github here
                        </p>
                        <div className="btn-actions">
                            <a href="https://github.com/haroldhang" target="blank" className="btn-primary btn-primary--stand">Visit my Github</a>
                        </div>
                    </div>

                </section>
            </main>
            <Footer/>
        </div>
    )

}

export default Home
