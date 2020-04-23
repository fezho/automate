import { Action } from '@ngrx/store';

export interface NodesAction extends Action {
  payload?: any;
}

export const LIST_NODES_SUCCESS = 'LIST_NODES_SUCCESS';
export const listNodesSuccess = (payload): NodesAction => ({ type: LIST_NODES_SUCCESS, payload });

export const LIST_NODES = 'LIST_NODES';
export const listNodes = (payload): NodesAction => ({ type: LIST_NODES, payload });




