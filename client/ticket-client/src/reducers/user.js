import { SIGNIN, SIGNUP } from '../constants/actionTypes';

const userReducer = ( state = { userData: null}, action) =>{
    switch(action.type){
        case SIGNIN:
            return state;
        case SIGNUP:
            return state;
        default:
            return state;
    }
}

export default userReducer;