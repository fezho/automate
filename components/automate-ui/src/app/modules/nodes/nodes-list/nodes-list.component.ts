import { Component, OnInit } from '@angular/core';
import { Observable, Subject, timer } from 'rxjs';
import { Store } from '@ngrx/store';
import { NgrxStateAtom } from 'app/ngrx.reducers';

import { takeUntil } from 'rxjs/operators';

import * as actions from 'app/entities/nodes/nodes.actions';
import * as selectors from 'app/entities/nodes/nodes.selectors';
import {  Node} from 'app/entities/nodes/nodes.requests';

@Component({  
  selector: 'app-nodes-list',
  templateUrl: './nodes-list.component.html',
  styleUrls: ['./nodes-list.component.scss']
})

export class NodesListComponent implements OnInit {

  nodesList;

  public nodes$: Observable<Node[]>;
  private isDestroyed: Subject<boolean> = new Subject<boolean>();

  constructor(
    private store: Store<NgrxStateAtom>
  ) { }

  ngOnInit() {
    this.store.select(selectors.nodes).pipe(
      takeUntil(this.isDestroyed))
      .subscribe(nodesList => this.nodesList = nodesList);

    // poll for updated jobs data every 5 secs (arbitrary interval)
    timer(0, 5000).pipe(
      takeUntil(this.isDestroyed))
      .subscribe(() => {
        this.store.dispatch(actions.listNodes('{}'));
      });
    console.log("nodes", this.nodesList)

    // this.store.dispatch(new actions.ListNodes());

    // this.nodes$ = this.store.select(selectors.nodes);


    // this.http.get<RespNodes>(
    //   `${NODES_URL}`)
    //   .pipe(map(respNodes =>
    //     this.listNodesResp(respNodes)));

  }

  // listNodesResp(respListNodes: RespNodes): RespNodes {
  //     console.log("?", respListNodes)
  //   return {
  //     nodes: respListNodes.nodes
  //   };
  // }
}
