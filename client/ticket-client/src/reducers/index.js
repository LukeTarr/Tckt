import { combineReducers } from 'redux';

import tickets from './tickets';
import user from './user';

export default combineReducers({ tickets, user});