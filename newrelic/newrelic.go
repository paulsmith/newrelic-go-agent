// Go wrapper for the New Relic Agent SDK. Requires Linux and the SDK headers
// and libraries.
package newrelic

/*
#cgo LDFLAGS: -L/usr/local/lib -lnewrelic-collector-client -lnewrelic-common -lnewrelic-transaction
#include "newrelic_collector_client.h"
#include "newrelic_common.h"
#include "newrelic_transaction.h"
#include "stdlib.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

func nrError(i C.int, name string) error {
	if int(i) < -1 {
		return errors.New("newrelic: " + name)
	}
	return nil
}

func Init(license string, appName string, lang string, langVersion string) error {
	C.newrelic_register_message_handler((*[0]byte)(C.newrelic_message_handler))
	clicense := C.CString(license)
	defer C.free(unsafe.Pointer(clicense))
	cappName := C.CString(appName)
	defer C.free(unsafe.Pointer(cappName))
	clang := C.CString(lang)
	defer C.free(unsafe.Pointer(clang))
	clangVersion := C.CString(langVersion)
	defer C.free(unsafe.Pointer(clangVersion))
	rv := C.newrelic_init(clicense, cappName, clang, clangVersion)
	return nrError(rv, "initialize")
}

func RequestShutdown(reason string) error {
	ptr := C.CString(reason)
	defer C.free(unsafe.Pointer(ptr))
	rv := C.newrelic_request_shutdown(ptr)
	return nrError(rv, "request shutdown")
}

func BeginTransaction() int64 {
	id := C.newrelic_transaction_begin()
	return int64(id)
}

func SetTransactionName(txnID int64, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	rv := C.newrelic_transaction_set_name(C.long(txnID), cname)
	return nrError(rv, "set transaction name")
}

func BeginGenericSegment(txnID int64, parentID int64, name string) int64 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	id := C.newrelic_segment_generic_begin(C.long(txnID), C.long(parentID), cname)
	return int64(id)
}

func BeginDatastoreSegment(
	txnID int64,
	parentID int64,
	table string,
	operation string,
	sql string,
	rollupName string,
) int64 {
	ctable := C.CString(table)
	defer C.free(unsafe.Pointer(ctable))
	coperation := C.CString(operation)
	defer C.free(unsafe.Pointer(coperation))
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))
	crollupName := C.CString(rollupName)
	defer C.free(unsafe.Pointer(crollupName))
	id := C.newrelic_segment_datastore_begin(
		C.long(txnID),
		C.long(parentID),
		ctable,
		coperation,
		csql,
		crollupName,
		(*[0]byte)(C.newrelic_basic_literal_replacement_obfuscator),
	)
	return int64(id)
}

func EndSegment(txnID int64, parentID int64) error {
	rv := C.newrelic_segment_end(C.long(txnID), C.long(parentID))
	return nrError(rv, "end segment")
}

func SetTransactionRequestURL(txnID int64, url string) error {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	rv := C.newrelic_transaction_set_request_url(C.long(txnID), curl)
	return nrError(rv, "set transaction request url")
}

func EndTransaction(txnID int64) error {
	rv := C.newrelic_transaction_end(C.long(txnID))
	return nrError(rv, "end transaction")
}

func RecordMetric(name string, val float64) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	rv := C.newrelic_record_metric(cname, C.double(val))
	return nrError(rv, "record metric")
}
