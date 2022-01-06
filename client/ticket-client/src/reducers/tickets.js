import { CREATE } from '../constants/actionTypes';

const ticketReducer = ( state = { ticketData: null}, action) =>{
    switch(action.type){
        case CREATE:
            return state;
        default:
            return state;
    }
}

export default ticketReducer;