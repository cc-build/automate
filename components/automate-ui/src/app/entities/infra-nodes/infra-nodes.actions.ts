import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';

import { InfraNode } from './infra-nodes.model';

export enum NodeActionTypes {
  GET_ALL         = 'NODES::GET_ALL',
  GET_ALL_SUCCESS = 'NODES::GET_ALL::SUCCESS',
  GET_ALL_FAILURE = 'NODES::GET_ALL::FAILURE',
  GET             = 'NODES::GET',
  GET_SUCCESS     = 'NODES::GET::SUCCESS',
  GET_FAILURE     = 'NODES::GET::FAILURE',
}

export interface NodesSuccessPayload {
  nodes: InfraNode[];
  total: number;
}

export interface NodesPayload {
  nodeName: string;
  server_id: string;
  org_id: string;
  page: number;
  per_page: number;
}
export class GetNodes implements Action {
  readonly type = NodeActionTypes.GET_ALL;
  constructor(public payload: NodesPayload) { }
}
export class GetNodesSuccess implements Action {
  readonly type = NodeActionTypes.GET_ALL_SUCCESS;
  constructor(public payload: NodesSuccessPayload) { }
}
export class GetNodesFailure implements Action {
  readonly type = NodeActionTypes.GET_ALL_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class GetNode implements Action {
  readonly type = NodeActionTypes.GET;
  constructor(public payload: { server_id: string, org_id: string, name: string }) { }
}

export class GetNodeSuccess implements Action {
  readonly type = NodeActionTypes.GET_SUCCESS;
  constructor(public payload: InfraNode) { }
}

export class GetNodeFailure implements Action {
  readonly type = NodeActionTypes.GET_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export type NodeActions =
  | GetNodes
  | GetNodesSuccess
  | GetNodesFailure
  | GetNode
  | GetNodeSuccess
  | GetNodeFailure;
