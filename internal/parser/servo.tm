language servo(go);

eventBased = true
lang = "go"
package = "github.com/egoodhall/servo/internal/parser/parsegen"

:: lexer

%x initial,inMessageDefinition,inUnionDefinition,inServiceDefinition,inEnumDefinition,inOption;

ident = /[a-zA-Z_][a-zA-Z0-9_]+/

stringLiteral = /"([^"]|\\")*"/
intLiteral = /-?[0-9]+/
floatLiteral = /-?[0-9]+\.[0-9]+/
boolLiteral = /true|false/

<initial> {
  'enum': /enum/ { l.State = StateInEnumDefinition }
  'message': /message/ { l.State = StateInMessageDefinition }
  'union': /union/ { l.State = StateInUnionDefinition }
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
  Primitive: /string|int64|int32|float32|float64|bool/ 1
  Modifier: /[?!]/
  '[': /\[/
  ']': /\]/
  ':': /:/
}

<inServiceDefinition> {
  Name: /{ident}/ (class)
  'rpc': /rpc/
  '(': /\(/
  ')': /\)/
}

<inUnionDefinition> {
  Name: /{ident}/
}

<inEnumDefinition> {
  Name: /{ident}/
  Value: /[A-Z][A-Z0-9_]+/ 1
}

<inMessageDefinition,inUnionDefinition,inServiceDefinition> {
  ':': /:/
}

<inMessageDefinition,inUnionDefinition,inServiceDefinition,inEnumDefinition> {
  ';': /;/
  '{': /\{/
  '}': /\}/ { l.State = StateInitial }
}

<inOption> {
  Name: /({ident}\.)?{ident}/
  '=': /=/
  StringLiteral: /{stringLiteral}/
  BoolLiteral: /{boolLiteral}/ 1
  IntLiteral: /{intLiteral}/
  FloatLiteral: /{floatLiteral}/
  ';': /;/ { l.State = StateInitial }
}

:: parser

%input ServoFile;

# WhiteSpace
Comment: BlockComment | EolComment;
WhiteSpace: (Comment | WS)+;

OptionName -> OptionName: Name;
OptionString -> OptionString: StringLiteral;
OptionBool -> OptionBool: BoolLiteral;
OptionInt -> OptionInt: IntLiteral;
OptionFloat -> OptionFloat: FloatLiteral;
OptionValue: OptionString | OptionBool | OptionInt | OptionFloat;
Option: 'option' OptionName '=' OptionValue ';';

# Type Definitions
TypeRef: (Name | Primitive);
FieldName -> FieldName: Name;
ScalarType -> ScalarType: TypeRef;
MapKeyType -> MapKeyType: Primitive;
MapValueType -> MapValueType: TypeRef;
MapType: '[' MapKeyType ':' MapValueType ']';
ListElementType -> ListElement: TypeRef;
ListType: '[' ListElementType ']';
FieldMod -> FieldMod: Modifier;
FieldType: (ScalarType | MapType | ListType) FieldMod?;
FieldDef: (FieldName WhiteSpace? ':' WhiteSpace? FieldType WhiteSpace? ';') | error  (';'|'}');
FieldDefList: (FieldDef | FieldDef WhiteSpace);
MessageName -> MessageName: Name;
MessageDef: 'message' MessageName WhiteSpace? '{' WhiteSpace? FieldDefList* '}';

# Union Definitions
UnionMemberName -> FieldName: Name;
UnionMemberType -> ScalarType: TypeRef;
UnionMember: (UnionMemberName WhiteSpace? ':' WhiteSpace? UnionMemberType WhiteSpace? ';') | error  (';'|'}');
UnionMemberList: (UnionMember | UnionMember WhiteSpace);
UnionName -> UnionName: Name;
UnionDef: 'union' UnionName WhiteSpace? '{' WhiteSpace? UnionMemberList* '}';

# Service Definitions
RpcName -> RpcName: Name;
RpcRequest -> RpcRequest: TypeRef;
RpcResponse -> RpcResponse: TypeRef;
Method: 'rpc' RpcName WhiteSpace? '(' WhiteSpace? RpcRequest  WhiteSpace? ')'  (WhiteSpace? ':'  WhiteSpace? RpcResponse)? WhiteSpace? ';';
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
Definition: MessageDef | UnionDef | ServiceDef | EnumDef | Option | error (';'|'}');
DefinitionList: (Definition | Definition WhiteSpace);

#File Definition
ServoFile: WhiteSpace? DefinitionList*;
