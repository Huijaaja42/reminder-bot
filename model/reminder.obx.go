// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package model

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type reminder_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ReminderBinding = reminder_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 8669441415074368449,
}

// Reminder_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Reminder_ = struct {
	Id      *objectbox.PropertyUint64
	User    *objectbox.PropertyString
	Channel *objectbox.PropertyString
	Time    *objectbox.PropertyInt64
	Text    *objectbox.PropertyString
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ReminderBinding.Entity,
		},
	},
	User: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ReminderBinding.Entity,
		},
	},
	Channel: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &ReminderBinding.Entity,
		},
	},
	Time: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ReminderBinding.Entity,
		},
	},
	Text: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ReminderBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (reminder_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (reminder_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Reminder", 1, 8669441415074368449)
	model.Property("Id", 6, 1, 4499052059592951076)
	model.PropertyFlags(1)
	model.Property("User", 9, 2, 6867887414617279894)
	model.Property("Channel", 9, 3, 2274699697038449398)
	model.Property("Time", 6, 4, 6136756833444865304)
	model.Property("Text", 9, 5, 572957272024362721)
	model.EntityLastPropertyId(5, 572957272024362721)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (reminder_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Reminder).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (reminder_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Reminder).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (reminder_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (reminder_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Reminder)
	var offsetUser = fbutils.CreateStringOffset(fbb, obj.User)
	var offsetChannel = fbutils.CreateStringOffset(fbb, obj.Channel)
	var offsetText = fbutils.CreateStringOffset(fbb, obj.Text)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetUser)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetChannel)
	fbutils.SetInt64Slot(fbb, 3, obj.Time)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetText)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (reminder_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Reminder' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	return &Reminder{
		Id:      propId,
		User:    fbutils.GetStringSlot(table, 6),
		Channel: fbutils.GetStringSlot(table, 8),
		Time:    fbutils.GetInt64Slot(table, 10),
		Text:    fbutils.GetStringSlot(table, 12),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (reminder_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Reminder, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (reminder_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Reminder), nil)
	}
	return append(slice.([]*Reminder), object.(*Reminder))
}

// Box provides CRUD access to Reminder objects
type ReminderBox struct {
	*objectbox.Box
}

// BoxForReminder opens a box of Reminder objects
func BoxForReminder(ob *objectbox.ObjectBox) *ReminderBox {
	return &ReminderBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Reminder.Id property on the passed object will be assigned the new ID as well.
func (box *ReminderBox) Put(object *Reminder) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Reminder.Id property on the passed object will be assigned the new ID as well.
func (box *ReminderBox) Insert(object *Reminder) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *ReminderBox) Update(object *Reminder) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *ReminderBox) PutAsync(object *Reminder) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Reminder.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Reminder.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ReminderBox) PutMany(objects []*Reminder) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ReminderBox) Get(id uint64) (*Reminder, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Reminder), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *ReminderBox) GetMany(ids ...uint64) ([]*Reminder, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Reminder), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *ReminderBox) GetManyExisting(ids ...uint64) ([]*Reminder, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Reminder), nil
}

// GetAll reads all stored objects
func (box *ReminderBox) GetAll() ([]*Reminder, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reminder), nil
}

// Remove deletes a single object
func (box *ReminderBox) Remove(object *Reminder) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *ReminderBox) RemoveMany(objects ...*Reminder) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Reminder_ struct to create conditions.
// Keep the *ReminderQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ReminderBox) Query(conditions ...objectbox.Condition) *ReminderQuery {
	return &ReminderQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Reminder_ struct to create conditions.
// Keep the *ReminderQuery if you intend to execute the query multiple times.
func (box *ReminderBox) QueryOrError(conditions ...objectbox.Condition) (*ReminderQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ReminderQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See ReminderAsyncBox for more information.
func (box *ReminderBox) Async() *ReminderAsyncBox {
	return &ReminderAsyncBox{AsyncBox: box.Box.Async()}
}

// ReminderAsyncBox provides asynchronous operations on Reminder objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type ReminderAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForReminder creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use ReminderBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForReminder(ob *objectbox.ObjectBox, timeoutMs uint64) *ReminderAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &ReminderAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *ReminderAsyncBox) Put(object *Reminder) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *ReminderAsyncBox) Insert(object *Reminder) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *ReminderAsyncBox) Update(object *Reminder) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *ReminderAsyncBox) Remove(object *Reminder) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Reminder which Id is either 42 or 47:
//
//	box.Query(Reminder_.Id.In(42, 47)).Find()
type ReminderQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ReminderQuery) Find() ([]*Reminder, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reminder), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ReminderQuery) Offset(offset uint64) *ReminderQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ReminderQuery) Limit(limit uint64) *ReminderQuery {
	query.Query.Limit(limit)
	return query
}
