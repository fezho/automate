import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Nodes } from './nodes.model';

import { environment } from '../../../environments/environment';
const NODES_URL = environment.nodes_url;

export interface RespNodes {
  nodes: Node[];
}

export interface Node {
  id: string,
  name: string
}

@Injectable()
export class NodesRequests {

  constructor(private http: HttpClient) { }

  public listNodes(): Observable<Nodes> {
    console.log("called listNodes")
    return this.http.get<RespNodes>(
      `${NODES_URL}/search`)
      .pipe(map(respNodes =>
        this.listNodesResp(respNodes)));
  }

  private listNodesResp(respListNodes: RespNodes): RespNodes {
      console.log("?", respListNodes)
    return {
      nodes: respListNodes.nodes
    };
  }
  
}