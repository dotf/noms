// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_crunchbase_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Company",
			[]types.Field{
				types.Field{"Permalink", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Name", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"HomepageUrl", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"CategoryList", types.MakeCompoundType(types.SetKind, types.MakePrimitiveType(types.StringKind)), false},
				types.Field{"Market", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundingTotalUsd", types.MakePrimitiveType(types.Float64Kind), false},
				types.Field{"Status", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"CountryCode", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"StateCode", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Region", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"City", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundingRounds", types.MakePrimitiveType(types.UInt16Kind), false},
				types.Field{"FoundedAt", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"FirstFundingAt", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"LastFundingAt", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"Rounds", types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(ref.Ref{}, 1))), false},
			},
			types.Choices{},
		),
		types.MakeStructType("Round",
			[]types.Field{
				types.Field{"CompanyPermalink", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundingRoundPermalink", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundingRoundType", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundingRoundCode", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"FundedAt", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"RaisedAmountUsd", types.MakePrimitiveType(types.Float64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__mainPackageInFile_crunchbase_CachedRef = types.RegisterPackage(&p)
}

// Company

type Company struct {
	_Permalink       string
	_Name            string
	_HomepageUrl     string
	_CategoryList    SetOfString
	_Market          string
	_FundingTotalUsd float64
	_Status          string
	_CountryCode     string
	_StateCode       string
	_Region          string
	_City            string
	_FundingRounds   uint16
	_FoundedAt       int64
	_FirstFundingAt  int64
	_LastFundingAt   int64
	_Rounds          SetOfRefOfRound

	ref *ref.Ref
}

func NewCompany() Company {
	return Company{
		_Permalink:       "",
		_Name:            "",
		_HomepageUrl:     "",
		_CategoryList:    NewSetOfString(),
		_Market:          "",
		_FundingTotalUsd: float64(0),
		_Status:          "",
		_CountryCode:     "",
		_StateCode:       "",
		_Region:          "",
		_City:            "",
		_FundingRounds:   uint16(0),
		_FoundedAt:       int64(0),
		_FirstFundingAt:  int64(0),
		_LastFundingAt:   int64(0),
		_Rounds:          NewSetOfRefOfRound(),

		ref: &ref.Ref{},
	}
}

type CompanyDef struct {
	Permalink       string
	Name            string
	HomepageUrl     string
	CategoryList    SetOfStringDef
	Market          string
	FundingTotalUsd float64
	Status          string
	CountryCode     string
	StateCode       string
	Region          string
	City            string
	FundingRounds   uint16
	FoundedAt       int64
	FirstFundingAt  int64
	LastFundingAt   int64
	Rounds          SetOfRefOfRoundDef
}

func (def CompanyDef) New() Company {
	return Company{
		_Permalink:       def.Permalink,
		_Name:            def.Name,
		_HomepageUrl:     def.HomepageUrl,
		_CategoryList:    def.CategoryList.New(),
		_Market:          def.Market,
		_FundingTotalUsd: def.FundingTotalUsd,
		_Status:          def.Status,
		_CountryCode:     def.CountryCode,
		_StateCode:       def.StateCode,
		_Region:          def.Region,
		_City:            def.City,
		_FundingRounds:   def.FundingRounds,
		_FoundedAt:       def.FoundedAt,
		_FirstFundingAt:  def.FirstFundingAt,
		_LastFundingAt:   def.LastFundingAt,
		_Rounds:          def.Rounds.New(),
		ref:              &ref.Ref{},
	}
}

func (s Company) Def() (d CompanyDef) {
	d.Permalink = s._Permalink
	d.Name = s._Name
	d.HomepageUrl = s._HomepageUrl
	d.CategoryList = s._CategoryList.Def()
	d.Market = s._Market
	d.FundingTotalUsd = s._FundingTotalUsd
	d.Status = s._Status
	d.CountryCode = s._CountryCode
	d.StateCode = s._StateCode
	d.Region = s._Region
	d.City = s._City
	d.FundingRounds = s._FundingRounds
	d.FoundedAt = s._FoundedAt
	d.FirstFundingAt = s._FirstFundingAt
	d.LastFundingAt = s._LastFundingAt
	d.Rounds = s._Rounds.Def()
	return
}

var __typeForCompany types.Type

func (m Company) Type() types.Type {
	return __typeForCompany
}

func init() {
	__typeForCompany = types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 0)
	types.RegisterStruct(__typeForCompany, builderForCompany, readerForCompany)
}

func builderForCompany(values []types.Value) types.Value {
	i := 0
	s := Company{ref: &ref.Ref{}}
	s._Permalink = values[i].(types.String).String()
	i++
	s._Name = values[i].(types.String).String()
	i++
	s._HomepageUrl = values[i].(types.String).String()
	i++
	s._CategoryList = values[i].(SetOfString)
	i++
	s._Market = values[i].(types.String).String()
	i++
	s._FundingTotalUsd = float64(values[i].(types.Float64))
	i++
	s._Status = values[i].(types.String).String()
	i++
	s._CountryCode = values[i].(types.String).String()
	i++
	s._StateCode = values[i].(types.String).String()
	i++
	s._Region = values[i].(types.String).String()
	i++
	s._City = values[i].(types.String).String()
	i++
	s._FundingRounds = uint16(values[i].(types.UInt16))
	i++
	s._FoundedAt = int64(values[i].(types.Int64))
	i++
	s._FirstFundingAt = int64(values[i].(types.Int64))
	i++
	s._LastFundingAt = int64(values[i].(types.Int64))
	i++
	s._Rounds = values[i].(SetOfRefOfRound)
	i++
	return s
}

func readerForCompany(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Company)
	values = append(values, types.NewString(s._Permalink))
	values = append(values, types.NewString(s._Name))
	values = append(values, types.NewString(s._HomepageUrl))
	values = append(values, s._CategoryList)
	values = append(values, types.NewString(s._Market))
	values = append(values, types.Float64(s._FundingTotalUsd))
	values = append(values, types.NewString(s._Status))
	values = append(values, types.NewString(s._CountryCode))
	values = append(values, types.NewString(s._StateCode))
	values = append(values, types.NewString(s._Region))
	values = append(values, types.NewString(s._City))
	values = append(values, types.UInt16(s._FundingRounds))
	values = append(values, types.Int64(s._FoundedAt))
	values = append(values, types.Int64(s._FirstFundingAt))
	values = append(values, types.Int64(s._LastFundingAt))
	values = append(values, s._Rounds)
	return values
}

func (s Company) Equals(other types.Value) bool {
	return other != nil && __typeForCompany.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Company) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Company) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForCompany.Chunks()...)
	chunks = append(chunks, s._CategoryList.Chunks()...)
	chunks = append(chunks, s._Rounds.Chunks()...)
	return
}

func (s Company) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._Permalink))
	ret = append(ret, types.NewString(s._Name))
	ret = append(ret, types.NewString(s._HomepageUrl))
	ret = append(ret, s._CategoryList)
	ret = append(ret, types.NewString(s._Market))
	ret = append(ret, types.Float64(s._FundingTotalUsd))
	ret = append(ret, types.NewString(s._Status))
	ret = append(ret, types.NewString(s._CountryCode))
	ret = append(ret, types.NewString(s._StateCode))
	ret = append(ret, types.NewString(s._Region))
	ret = append(ret, types.NewString(s._City))
	ret = append(ret, types.UInt16(s._FundingRounds))
	ret = append(ret, types.Int64(s._FoundedAt))
	ret = append(ret, types.Int64(s._FirstFundingAt))
	ret = append(ret, types.Int64(s._LastFundingAt))
	ret = append(ret, s._Rounds)
	return
}

func (s Company) Permalink() string {
	return s._Permalink
}

func (s Company) SetPermalink(val string) Company {
	s._Permalink = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) Name() string {
	return s._Name
}

func (s Company) SetName(val string) Company {
	s._Name = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) HomepageUrl() string {
	return s._HomepageUrl
}

func (s Company) SetHomepageUrl(val string) Company {
	s._HomepageUrl = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) CategoryList() SetOfString {
	return s._CategoryList
}

func (s Company) SetCategoryList(val SetOfString) Company {
	s._CategoryList = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) Market() string {
	return s._Market
}

func (s Company) SetMarket(val string) Company {
	s._Market = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) FundingTotalUsd() float64 {
	return s._FundingTotalUsd
}

func (s Company) SetFundingTotalUsd(val float64) Company {
	s._FundingTotalUsd = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) Status() string {
	return s._Status
}

func (s Company) SetStatus(val string) Company {
	s._Status = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) CountryCode() string {
	return s._CountryCode
}

func (s Company) SetCountryCode(val string) Company {
	s._CountryCode = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) StateCode() string {
	return s._StateCode
}

func (s Company) SetStateCode(val string) Company {
	s._StateCode = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) Region() string {
	return s._Region
}

func (s Company) SetRegion(val string) Company {
	s._Region = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) City() string {
	return s._City
}

func (s Company) SetCity(val string) Company {
	s._City = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) FundingRounds() uint16 {
	return s._FundingRounds
}

func (s Company) SetFundingRounds(val uint16) Company {
	s._FundingRounds = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) FoundedAt() int64 {
	return s._FoundedAt
}

func (s Company) SetFoundedAt(val int64) Company {
	s._FoundedAt = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) FirstFundingAt() int64 {
	return s._FirstFundingAt
}

func (s Company) SetFirstFundingAt(val int64) Company {
	s._FirstFundingAt = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) LastFundingAt() int64 {
	return s._LastFundingAt
}

func (s Company) SetLastFundingAt(val int64) Company {
	s._LastFundingAt = val
	s.ref = &ref.Ref{}
	return s
}

func (s Company) Rounds() SetOfRefOfRound {
	return s._Rounds
}

func (s Company) SetRounds(val SetOfRefOfRound) Company {
	s._Rounds = val
	s.ref = &ref.Ref{}
	return s
}

// Round

type Round struct {
	_CompanyPermalink      string
	_FundingRoundPermalink string
	_FundingRoundType      string
	_FundingRoundCode      string
	_FundedAt              int64
	_RaisedAmountUsd       float64

	ref *ref.Ref
}

func NewRound() Round {
	return Round{
		_CompanyPermalink:      "",
		_FundingRoundPermalink: "",
		_FundingRoundType:      "",
		_FundingRoundCode:      "",
		_FundedAt:              int64(0),
		_RaisedAmountUsd:       float64(0),

		ref: &ref.Ref{},
	}
}

type RoundDef struct {
	CompanyPermalink      string
	FundingRoundPermalink string
	FundingRoundType      string
	FundingRoundCode      string
	FundedAt              int64
	RaisedAmountUsd       float64
}

func (def RoundDef) New() Round {
	return Round{
		_CompanyPermalink:      def.CompanyPermalink,
		_FundingRoundPermalink: def.FundingRoundPermalink,
		_FundingRoundType:      def.FundingRoundType,
		_FundingRoundCode:      def.FundingRoundCode,
		_FundedAt:              def.FundedAt,
		_RaisedAmountUsd:       def.RaisedAmountUsd,
		ref:                    &ref.Ref{},
	}
}

func (s Round) Def() (d RoundDef) {
	d.CompanyPermalink = s._CompanyPermalink
	d.FundingRoundPermalink = s._FundingRoundPermalink
	d.FundingRoundType = s._FundingRoundType
	d.FundingRoundCode = s._FundingRoundCode
	d.FundedAt = s._FundedAt
	d.RaisedAmountUsd = s._RaisedAmountUsd
	return
}

var __typeForRound types.Type

func (m Round) Type() types.Type {
	return __typeForRound
}

func init() {
	__typeForRound = types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 1)
	types.RegisterStruct(__typeForRound, builderForRound, readerForRound)
}

func builderForRound(values []types.Value) types.Value {
	i := 0
	s := Round{ref: &ref.Ref{}}
	s._CompanyPermalink = values[i].(types.String).String()
	i++
	s._FundingRoundPermalink = values[i].(types.String).String()
	i++
	s._FundingRoundType = values[i].(types.String).String()
	i++
	s._FundingRoundCode = values[i].(types.String).String()
	i++
	s._FundedAt = int64(values[i].(types.Int64))
	i++
	s._RaisedAmountUsd = float64(values[i].(types.Float64))
	i++
	return s
}

func readerForRound(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Round)
	values = append(values, types.NewString(s._CompanyPermalink))
	values = append(values, types.NewString(s._FundingRoundPermalink))
	values = append(values, types.NewString(s._FundingRoundType))
	values = append(values, types.NewString(s._FundingRoundCode))
	values = append(values, types.Int64(s._FundedAt))
	values = append(values, types.Float64(s._RaisedAmountUsd))
	return values
}

func (s Round) Equals(other types.Value) bool {
	return other != nil && __typeForRound.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Round) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Round) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForRound.Chunks()...)
	return
}

func (s Round) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._CompanyPermalink))
	ret = append(ret, types.NewString(s._FundingRoundPermalink))
	ret = append(ret, types.NewString(s._FundingRoundType))
	ret = append(ret, types.NewString(s._FundingRoundCode))
	ret = append(ret, types.Int64(s._FundedAt))
	ret = append(ret, types.Float64(s._RaisedAmountUsd))
	return
}

func (s Round) CompanyPermalink() string {
	return s._CompanyPermalink
}

func (s Round) SetCompanyPermalink(val string) Round {
	s._CompanyPermalink = val
	s.ref = &ref.Ref{}
	return s
}

func (s Round) FundingRoundPermalink() string {
	return s._FundingRoundPermalink
}

func (s Round) SetFundingRoundPermalink(val string) Round {
	s._FundingRoundPermalink = val
	s.ref = &ref.Ref{}
	return s
}

func (s Round) FundingRoundType() string {
	return s._FundingRoundType
}

func (s Round) SetFundingRoundType(val string) Round {
	s._FundingRoundType = val
	s.ref = &ref.Ref{}
	return s
}

func (s Round) FundingRoundCode() string {
	return s._FundingRoundCode
}

func (s Round) SetFundingRoundCode(val string) Round {
	s._FundingRoundCode = val
	s.ref = &ref.Ref{}
	return s
}

func (s Round) FundedAt() int64 {
	return s._FundedAt
}

func (s Round) SetFundedAt(val int64) Round {
	s._FundedAt = val
	s.ref = &ref.Ref{}
	return s
}

func (s Round) RaisedAmountUsd() float64 {
	return s._RaisedAmountUsd
}

func (s Round) SetRaisedAmountUsd(val float64) Round {
	s._RaisedAmountUsd = val
	s.ref = &ref.Ref{}
	return s
}

// MapOfStringToRefOfCompany

type MapOfStringToRefOfCompany struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToRefOfCompany() MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{types.NewTypedMap(__typeForMapOfStringToRefOfCompany), &ref.Ref{}}
}

type MapOfStringToRefOfCompanyDef map[string]ref.Ref

func (def MapOfStringToRefOfCompanyDef) New() MapOfStringToRefOfCompany {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfCompany(v))
	}
	return MapOfStringToRefOfCompany{types.NewTypedMap(__typeForMapOfStringToRefOfCompany, kv...), &ref.Ref{}}
}

func (m MapOfStringToRefOfCompany) Def() MapOfStringToRefOfCompanyDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(RefOfCompany).TargetRef()
		return false
	})
	return def
}

func (m MapOfStringToRefOfCompany) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfStringToRefOfCompany.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToRefOfCompany) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfCompany) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToRefOfCompany) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToRefOfCompany.
var __typeForMapOfStringToRefOfCompany types.Type

func (m MapOfStringToRefOfCompany) Type() types.Type {
	return __typeForMapOfStringToRefOfCompany
}

func init() {
	__typeForMapOfStringToRefOfCompany = types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 0)))
	types.RegisterValue(__typeForMapOfStringToRefOfCompany, builderForMapOfStringToRefOfCompany, readerForMapOfStringToRefOfCompany)
}

func builderForMapOfStringToRefOfCompany(v types.Value) types.Value {
	return MapOfStringToRefOfCompany{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfStringToRefOfCompany(v types.Value) types.Value {
	return v.(MapOfStringToRefOfCompany).m
}

func (m MapOfStringToRefOfCompany) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfCompany) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfCompany) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfCompany) Get(p string) RefOfCompany {
	return m.m.Get(types.NewString(p)).(RefOfCompany)
}

func (m MapOfStringToRefOfCompany) MaybeGet(p string) (RefOfCompany, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfCompany(ref.Ref{}), false
	}
	return v.(RefOfCompany), ok
}

func (m MapOfStringToRefOfCompany) Set(k string, v RefOfCompany) MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfCompany) Remove(p string) MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToRefOfCompanyIterCallback func(k string, v RefOfCompany) (stop bool)

func (m MapOfStringToRefOfCompany) Iter(cb MapOfStringToRefOfCompanyIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

type MapOfStringToRefOfCompanyIterAllCallback func(k string, v RefOfCompany)

func (m MapOfStringToRefOfCompany) IterAll(cb MapOfStringToRefOfCompanyIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

func (m MapOfStringToRefOfCompany) IterAllP(concurrency int, cb MapOfStringToRefOfCompanyIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

type MapOfStringToRefOfCompanyFilterCallback func(k string, v RefOfCompany) (keep bool)

func (m MapOfStringToRefOfCompany) Filter(cb MapOfStringToRefOfCompanyFilterCallback) MapOfStringToRefOfCompany {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCompany))
	})
	return MapOfStringToRefOfCompany{out, &ref.Ref{}}
}

// SetOfString

type SetOfString struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfString() SetOfString {
	return SetOfString{types.NewTypedSet(__typeForSetOfString), &ref.Ref{}}
}

type SetOfStringDef map[string]bool

func (def SetOfStringDef) New() SetOfString {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.NewString(d)
		i++
	}
	return SetOfString{types.NewTypedSet(__typeForSetOfString, l...), &ref.Ref{}}
}

func (s SetOfString) Def() SetOfStringDef {
	def := make(map[string]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(types.String).String()] = true
		return false
	})
	return def
}

func (s SetOfString) Equals(other types.Value) bool {
	return other != nil && __typeForSetOfString.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SetOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.Type().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

func (s SetOfString) ChildValues() []types.Value {
	return append([]types.Value{}, s.s.ChildValues()...)
}

// A Noms Value that describes SetOfString.
var __typeForSetOfString types.Type

func (m SetOfString) Type() types.Type {
	return __typeForSetOfString
}

func init() {
	__typeForSetOfString = types.MakeCompoundType(types.SetKind, types.MakePrimitiveType(types.StringKind))
	types.RegisterValue(__typeForSetOfString, builderForSetOfString, readerForSetOfString)
}

func builderForSetOfString(v types.Value) types.Value {
	return SetOfString{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfString(v types.Value) types.Value {
	return v.(SetOfString).s
}

func (s SetOfString) Empty() bool {
	return s.s.Empty()
}

func (s SetOfString) Len() uint64 {
	return s.s.Len()
}

func (s SetOfString) Has(p string) bool {
	return s.s.Has(types.NewString(p))
}

type SetOfStringIterCallback func(p string) (stop bool)

func (s SetOfString) Iter(cb SetOfStringIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(types.String).String())
	})
}

type SetOfStringIterAllCallback func(p string)

func (s SetOfString) IterAll(cb SetOfStringIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(types.String).String())
	})
}

func (s SetOfString) IterAllP(concurrency int, cb SetOfStringIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(types.String).String())
	})
}

type SetOfStringFilterCallback func(p string) (keep bool)

func (s SetOfString) Filter(cb SetOfStringFilterCallback) SetOfString {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(types.String).String())
	})
	return SetOfString{out, &ref.Ref{}}
}

func (s SetOfString) Insert(p ...string) SetOfString {
	return SetOfString{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfString) Remove(p ...string) SetOfString {
	return SetOfString{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfString) Union(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfString) Subtract(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfString) Any() string {
	return s.s.Any().(types.String).String()
}

func (s SetOfString) fromStructSlice(p []SetOfString) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfString) fromElemSlice(p []string) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.NewString(v)
	}
	return r
}

// SetOfRefOfRound

type SetOfRefOfRound struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRefOfRound() SetOfRefOfRound {
	return SetOfRefOfRound{types.NewTypedSet(__typeForSetOfRefOfRound), &ref.Ref{}}
}

type SetOfRefOfRoundDef map[ref.Ref]bool

func (def SetOfRefOfRoundDef) New() SetOfRefOfRound {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfRound(d)
		i++
	}
	return SetOfRefOfRound{types.NewTypedSet(__typeForSetOfRefOfRound, l...), &ref.Ref{}}
}

func (s SetOfRefOfRound) Def() SetOfRefOfRoundDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfRound).TargetRef()] = true
		return false
	})
	return def
}

func (s SetOfRefOfRound) Equals(other types.Value) bool {
	return other != nil && __typeForSetOfRefOfRound.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SetOfRefOfRound) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfRound) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.Type().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

func (s SetOfRefOfRound) ChildValues() []types.Value {
	return append([]types.Value{}, s.s.ChildValues()...)
}

// A Noms Value that describes SetOfRefOfRound.
var __typeForSetOfRefOfRound types.Type

func (m SetOfRefOfRound) Type() types.Type {
	return __typeForSetOfRefOfRound
}

func init() {
	__typeForSetOfRefOfRound = types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 1)))
	types.RegisterValue(__typeForSetOfRefOfRound, builderForSetOfRefOfRound, readerForSetOfRefOfRound)
}

func builderForSetOfRefOfRound(v types.Value) types.Value {
	return SetOfRefOfRound{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfRefOfRound(v types.Value) types.Value {
	return v.(SetOfRefOfRound).s
}

func (s SetOfRefOfRound) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfRound) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfRound) Has(p RefOfRound) bool {
	return s.s.Has(p)
}

type SetOfRefOfRoundIterCallback func(p RefOfRound) (stop bool)

func (s SetOfRefOfRound) Iter(cb SetOfRefOfRoundIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfRound))
	})
}

type SetOfRefOfRoundIterAllCallback func(p RefOfRound)

func (s SetOfRefOfRound) IterAll(cb SetOfRefOfRoundIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfRound))
	})
}

func (s SetOfRefOfRound) IterAllP(concurrency int, cb SetOfRefOfRoundIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RefOfRound))
	})
}

type SetOfRefOfRoundFilterCallback func(p RefOfRound) (keep bool)

func (s SetOfRefOfRound) Filter(cb SetOfRefOfRoundFilterCallback) SetOfRefOfRound {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(RefOfRound))
	})
	return SetOfRefOfRound{out, &ref.Ref{}}
}

func (s SetOfRefOfRound) Insert(p ...RefOfRound) SetOfRefOfRound {
	return SetOfRefOfRound{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRound) Remove(p ...RefOfRound) SetOfRefOfRound {
	return SetOfRefOfRound{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRound) Union(others ...SetOfRefOfRound) SetOfRefOfRound {
	return SetOfRefOfRound{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRound) Subtract(others ...SetOfRefOfRound) SetOfRefOfRound {
	return SetOfRefOfRound{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRound) Any() RefOfRound {
	return s.s.Any().(RefOfRound)
}

func (s SetOfRefOfRound) fromStructSlice(p []SetOfRefOfRound) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfRound) fromElemSlice(p []RefOfRound) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfCompany

type RefOfCompany struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfCompany(target ref.Ref) RefOfCompany {
	return RefOfCompany{target, &ref.Ref{}}
}

func (r RefOfCompany) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfCompany) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfCompany) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfCompany.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfCompany) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfCompany) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfCompany.
var __typeForRefOfCompany types.Type

func (m RefOfCompany) Type() types.Type {
	return __typeForRefOfCompany
}

func init() {
	__typeForRefOfCompany = types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 0))
	types.RegisterRef(__typeForRefOfCompany, builderForRefOfCompany)
}

func builderForRefOfCompany(r ref.Ref) types.Value {
	return NewRefOfCompany(r)
}

func (r RefOfCompany) TargetValue(cs chunks.ChunkSource) Company {
	return types.ReadValue(r.target, cs).(Company)
}

func (r RefOfCompany) SetTargetValue(val Company, cs chunks.ChunkSink) RefOfCompany {
	return NewRefOfCompany(types.WriteValue(val, cs))
}

// RefOfRound

type RefOfRound struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfRound(target ref.Ref) RefOfRound {
	return RefOfRound{target, &ref.Ref{}}
}

func (r RefOfRound) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfRound) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfRound) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfRound.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfRound) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfRound) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfRound.
var __typeForRefOfRound types.Type

func (m RefOfRound) Type() types.Type {
	return __typeForRefOfRound
}

func init() {
	__typeForRefOfRound = types.MakeCompoundType(types.RefKind, types.MakeType(__mainPackageInFile_crunchbase_CachedRef, 1))
	types.RegisterRef(__typeForRefOfRound, builderForRefOfRound)
}

func builderForRefOfRound(r ref.Ref) types.Value {
	return NewRefOfRound(r)
}

func (r RefOfRound) TargetValue(cs chunks.ChunkSource) Round {
	return types.ReadValue(r.target, cs).(Round)
}

func (r RefOfRound) SetTargetValue(val Round, cs chunks.ChunkSink) RefOfRound {
	return NewRefOfRound(types.WriteValue(val, cs))
}