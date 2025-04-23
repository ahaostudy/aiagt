namespace go workflowsvc

include './base.thrift'
include './types.thrift'

struct Workflow {
    1: required i64 id
    2: required string name
    3: required string description
    4: required string description_md
    5: required i64 author_id
    6: required bool is_private
    7: required string logo
    8: optional base.Time published_at
    9: required base.Time created_at
    10: required base.Time updated_at
}

struct WorkflowNode {
    1: required i64 id
    2: required string name
    3: required string type
    4: required types.any input_mapper
    5: optional types.any output_schema
    6: optional types.any batch_field
    7: required i64 workflow_id
    8: required list<i64> next_ids
    9: required types.any node_params
    10: required base.Time created_at
    11: required base.Time updated_at
}

struct CreateWorkflowReq {
    1: required string name
    2: required string description
    3: required string description_md
    4: required i64 author_id
    5: required bool is_private
    6: required string logo
}

struct UpdateWorkflowReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional string name
    3: optional string description
    4: optional string description_md
    5: optional bool is_private
    6: optional string logo
    7: optional base.Time published_at
}

struct CreateWorkflowNodeReq {
    1: required string name
    2: required string type
    3: required types.any input_mapper
    4: optional types.any output_schema
    5: optional types.any batch_field
    6: required i64 workflow_id
    7: required list<i64> next_ids
    8: required types.any node_params
}

struct UpdateWorkflowNodeReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional string name
    3: optional string type
    4: optional types.any input_mapper
    5: optional types.any output_schema
    6: optional types.any batch_field
    7: optional list<i64> next_ids
    8: optional types.any node_params
}

struct CallWorkflowReq {
    1: required i64 workflow_id
    2: required string request
}

struct CallWorkflowResp {
    1: required string response
}

service WorkflowService {
    base.Empty CreateWorkflow(1: CreateWorkflowReq req)
    base.Empty UpdateWorkflow(1: UpdateWorkflowReq req)
    base.Empty DeleteWorkflow(1: base.IDReq req)
    Workflow GetWorkflowByID(1: base.IDReq req)

    base.Empty CreateWorkflowNode(1: CreateWorkflowNodeReq req)
    base.Empty UpdateWorkflowNode(1: UpdateWorkflowNodeReq req)
    base.Empty DeleteWorkflowNode(1: base.IDReq req)
    WorkflowNode GetWorkflowNodeByID(1: base.IDReq req)

    CallWorkflowResp CallWorkflow(1: CallWorkflowReq req)
}
