import { useState } from "react";
import { signin, signup } from "../../actions/user";
import { useDispatch } from "react-redux";
const Auth = () => {
    const dispatch = useDispatch();
    const [isSignUp, setIsSignUp] = useState(false);
    const [formData, setFormData] = useState({
        email:"",
        password:"",
        passwordConfirm:"",
    });

    const clear = () => {
        setFormData({
            email:"",
            password:"",
            passwordConfirm:"",
        })
    }
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log(formData);
        if(isSignUp){
            console.log("Sign Up Submit");
            dispatch(signup(formData));
        }
        else{
            console.log("Sing In Submit");
            dispatch(signin(formData));
        }
        clear();
    }

    return(
        <div>
        <div>
            <form onSubmit={handleSubmit}>
                <label>
                    Email:
                    <input type="email" required={true} name="email" onChange={(e) => setFormData({ ...formData, email: e.target.value })}/>
                </label>
                <label>
                    Password:
                    <input type="password" required={true} name="password"  onChange={(e) => setFormData({ ...formData, password: e.target.value })}/>
                </label>
                {!isSignUp ? <></> : 
                <label>
                    Confirm Password:
                    <input type="password" required={true} name="passwordConfirm"  onChange={(e) => setFormData({ ...formData, passwordConfirm: e.target.value })}/>
                </label>
                }   
                <button type="submit">{isSignUp ? "Sign Up" : "Sign In"}</button>
            </form>
        </div>
        {
        <div>
            {isSignUp ? <h1>Already have an account?</h1>: <h1>Don't have an account?</h1>}
            <button onClick={() => setIsSignUp(!isSignUp)}>{isSignUp ? <h1>Sign In</h1>:<h1>Sign Up</h1>}</button>
        </div>
        }
        </div>
    );
}

export default Auth;