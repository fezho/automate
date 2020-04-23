import { createSelector} from '@ngrx/store';


export const nodesState = state => state.nodes;

export const nodes = createSelector(
  nodesState,
  (state) => state.nodes
);

export const nodesList = createSelector(nodesState, state => state.nodes);
