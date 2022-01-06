import { SIGNIN, SIGNUP } from "../constants/actionTypes";
import * as api from '../api';

export const signin = (formData) => async (dispatch) => {
    try {
        console.log("IN SINGIN ACTION, formData:", formData);
        const { data } = await api.signin(formData);
        console.log("SINGIN RETDATA: ", data);
        dispatch({type: SIGNIN, data});

    } catch (error) {
        console.log(error);
    }
}

export const signup = (formData) => async (dispatch) => {
    try {
        console.log("IN SIGNUP ACTION, formData:", formData);
        const { data } = await api.signup(formData);
        console.log("SINGUP RETDATA: ", data);
        dispatch({type: SIGNUP, data});

    } catch (error) {
        console.log(error);
    }
}
