language servo(go);

eventBased = true
lang = "go"
package = "github.com/egoodhall/servo/internal/parser/parsegen"

:: lexer

%x initial,inMessageDefinition,inServiceDefinition,inEnumDefinition,inOption;

ident = /[a-zA-Z_][a-zA-Z0-9_]+/

stringLiteral = /"([^"]|\\")*"/

<initial> {
  'enum': /enum/ { l.State = StateInEnumDefinition }
  'message': /message/ { l.State = StateInMessageDefinition }
  'service': /service/ { l.State = StateInServiceDefinition }
  'option': /option/ { l.State = StateInOption }
}

<*> {
  WS: /[ \t\n\r]+/ (space)
  EolComment: /\/\/[^\n]+\n/ (space)
  BlockComment: /\/\*([^*]|\*+[^*\/])*\**\*\// (space)
  error:
}

<inMessageDefinition> {
  Name: /{ident}/
  Primitive: /string|int64|int32|float32|float64/ 1
  Modifier: /[?!]/
}

<inServiceDefinition> {
  Name: /{ident}/ (class)
  'rpc': /rpc/
  'pub': /pub/
  '(': /\(/
  ')': /\)/
}

<inEnumDefinition> {
  Name: /{ident}/
  Value: /[A-Z][A-Z0-9_]+/ 1
}

<inMessageDefinition,inServiceDefinition> {
  ':': /:/
}

<inMessageDefinition,inServiceDefinition,inEnumDefinition> {
  ';': /;/
  '{': /\{/
  '}': /\}/ { l.State = StateInitial }
}

<inOption> {
  Name: /({ident}\.)?{ident}/
  '=': /=/
  StringLiteral: /{stringLiteral}/
  ';': /;/ { l.State = StateInitial }
}

:: parser

%input ServoFile;

# WhiteSpace
Comment: BlockComment | EolComment;
WhiteSpace: (Comment | WS)+;

OptionName -> OptionName: Name;
OptionValue -> OptionValue: StringLiteral;
Option: 'option' OptionName '=' OptionValue ';';

# Type Definitions
TypeRef: (Name | Primitive) | error (';'|'}');
FieldName -> FieldName: Name;
FieldType -> FieldType: TypeRef Modifier?;
FieldDef: (FieldName WhiteSpace? ':' WhiteSpace? FieldType WhiteSpace? ';') | error  (';'|'}');
FieldDefList: (FieldDef | FieldDef WhiteSpace);
MessageName -> MessageName: Name;
MessageDef: 'message' MessageName WhiteSpace? '{' WhiteSpace? FieldDefList* '}';

# Service Definitions
RpcName -> RpcName: Name;
RpcRequest -> RpcRequest: TypeRef;
RpcResponse -> RpcResponse: TypeRef;
RpcMethod: 'rpc' RpcName WhiteSpace? '(' WhiteSpace? RpcRequest  WhiteSpace? ')'  WhiteSpace? ':'  WhiteSpace? RpcResponse  WhiteSpace? ';';
PubName -> PubName: Name;
PubMessage -> PubMessage: TypeRef;
PubMethod: 'pub' PubName WhiteSpace? '(' WhiteSpace? PubMessage WhiteSpace? ')' WhiteSpace? ';';
Method: PubMethod | RpcMethod | error  (';'|'}');
 
MethodList: (Method | Method WhiteSpace);
ServiceName -> ServiceName: Name;
ServiceDef: 'service' ServiceName WhiteSpace? '{' WhiteSpace? MethodList+ '}';

# Enum Definitions
EnumName -> EnumName: Name;
EnumValue -> EnumValue: Value;
EnumField: EnumValue WhiteSpace? ';';
EnumValueList: (EnumField | EnumField WhiteSpace);
EnumDef: 'enum' EnumName WhiteSpace? '{' WhiteSpace? EnumValueList+ '}';

# Elements
Definition: MessageDef | ServiceDef | EnumDef | Option | error (';'|'}');
DefinitionList: (Definition | Definition WhiteSpace);

#File Definition
ServoFile: WhiteSpace? DefinitionList*;
