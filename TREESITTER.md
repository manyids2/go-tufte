# Treesitter API

```go

func Parse(content []byte, lang *Language) *Node
func ParseCtx(ctx context.Context, content []byte, lang *Language) (*Node
, error)
func NewParser() *Parser
func (p *Parser) SetLanguage(lang *Language)
func (p *Parser) Parse(oldTree *Tree, content []byte) *Tree
func (p *Parser) ParseCtx(ctx context.Context, oldTree *Tree, content []b
yte) (*Tree, error)
func (p *Parser) ParseInput(oldTree *Tree, input Input) *Tree
func (p *Parser) ParseInputCtx(ctx context.Context, oldTree *Tree, input
Input) (*Tree, error)
func (p *Parser) convertTSTree(ctx context.Context, tsTree *C.TSTree) (*T
ree, error)
func (p *Parser) OperationLimit() int
func (p *Parser) SetOperationLimit(limit int)
func (p *Parser) Reset()
func (p *Parser) SetIncludedRanges(ranges []Range)
func (p *Parser) Debug()
func (p *Parser) Close()
func (p *Parser) newTree(c *C.TSTree) *Tree
func (t *Tree) Copy() *Tree
func (t *Tree) RootNode() *Node
func (t *Tree) cachedNode(ptr C.TSNode) *Node
func (t *BaseTree) Close()
func (i EditInput) c() *C.TSInputEdit
func (t *Tree) Edit(i EditInput)
func NewLanguage(ptr unsafe.Pointer) *Language
func (l *Language) SymbolName(s Symbol) string
func (l *Language) SymbolType(s Symbol) SymbolType
func (l *Language) SymbolCount() uint32
func (l *Language) FieldName(idx int) string
func (t SymbolType) String() string
func (n Node) StartByte() uint32
func (n Node) EndByte() uint32
func (n Node) StartPoint() Point
func (n Node) EndPoint() Point
func (n Node) Symbol() Symbol
func (n Node) Type() string
func (n Node) String() string
func (n Node) Equal(other *Node) bool
func (n Node) IsNull() bool
func (n Node) IsNamed() bool
func (n Node) IsMissing() bool
func (n Node) IsExtra() bool
func (n Node) IsError() bool
func (n Node) HasChanges() bool
func (n Node) HasError() bool
func (n Node) Parent() *Node
func (n Node) Child(idx int) *Node
func (n Node) NamedChild(idx int) *Node
func (n Node) ChildCount() uint32
func (n Node) NamedChildCount() uint32
func (n Node) ChildByFieldName(name string) *Node
func (n Node) FieldNameForChild(idx int) string
func (n Node) NextSibling() *Node
func (n Node) NextNamedSibling() *Node
func (n Node) PrevSibling() *Node
func (n Node) PrevNamedSibling() *Node
func (n Node) Edit(i EditInput)
func (n Node) Content(input []byte) string
func (n Node) NamedDescendantForPointRange(start Point, end Point) *Node
func NewTreeCursor(n *Node) *TreeCursor
func (c *TreeCursor) Close()
func (c *TreeCursor) Reset(n *Node)
func (c *TreeCursor) CurrentNode() *Node
func (c *TreeCursor) CurrentFieldName() string
func (c *TreeCursor) GoToParent() bool
func (c *TreeCursor) GoToNextSibling() bool
func (c *TreeCursor) GoToFirstChild() bool
func (c *TreeCursor) GoToFirstChildForByte(b uint32) int64
func QueryErrorTypeToString(errorType QueryErrorType) string
func (qe *QueryError) Error() string
func NewQuery(pattern []byte, lang *Language) (*Query, error)
func (q *Query) Close()
func (q *Query) PatternCount() uint32
func (q *Query) CaptureCount() uint32
func (q *Query) StringCount() uint32
func (q *Query) PredicatesForPattern(patternIndex uint32) [][]QueryPredic
ateStep
func (q *Query) CaptureNameForId(id uint32) string
func (q *Query) StringValueForId(id uint32) string
func (q *Query) CaptureQuantifierForId(id uint32, captureId uint32) Quant
ifier
func NewQueryCursor() *QueryCursor
func (qc *QueryCursor) Exec(q *Query, n *Node)
func (qc *QueryCursor) SetPointRange(startPoint Point, endPoint Point)
func (qc *QueryCursor) Close()
func (qc *QueryCursor) NextMatch() (*QueryMatch, bool)
func (qc *QueryCursor) NextCapture() (*QueryMatch, uint32, bool)
func splitPredicates(steps []QueryPredicateStep) [][]QueryPredicateStep
func (qc *QueryCursor) FilterPredicates(m *QueryMatch, input []byte) *Que
ryMatch
func (m *readFuncsMap) register(f ReadFunc) int
func (m *readFuncsMap) unregister(id int)
func (m *readFuncsMap) get(id int) ReadFunc
func callReadFunc(id C.int, byteIndex C.uint32_t, position C.TSPoint, byt
esRead *C.uint32_t) *C.char
```

## Primitives

Choose nodes that do not have children, put them into json

```bash
# Start json
echo '[' > node-types-primitives.json

# Extract
jq '.[] | select(.children == null)' node-types.json \
  >> node-types-primitives.json

# Add commas, remove last comma
sed 's/^}/},/' -i node-types-primitives.json
sed '$ s/.$//' -i node-types-primitives.json

# End json
echo ']' >> node-types-primitives.json
```

## Elements

```bash
# Start json
echo '[' > node-types-elements.json

# Extract
jq '.[] | select(.children != null)' node-types.json \
  >> node-types-elements.json

# Add commas, remove last comma
sed 's/^}/},/' -i node-types-elements.json
sed '$ s/.$//' -i node-types-elements.json

# End json
echo ']' >> node-types-elements.json
```

## Blocks
