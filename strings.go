package maptools

import "strings"
import "sort"

// Returns a list of map keys for a map[string]string type
func StringStringMapKeys(v map[string]string)(out []string){
  out = make([]string, len(v))
  i := 0
  for k, _ := range(v) {
    out[i] = k
    i += 1
  }
  return
}

// Returns a list of map keys for a map[string][]string type
func StringStringsMapKeys(v map[string][]string)(out []string){
  out = make([]string, len(v))
  i := 0
  for k, _ := range(v) {
    out[i] = k
    i += 1
  }
  return
}

// Converts a map[string][]string to map[string][]string by performing 
// string.Join(, sep) on each value (and optionally sorting the values first.
func StringStringsJoin(v map[string][]string, sep string, sortVals bool)(out map[string]string){
  out = make(map[string]string)
  for  k, v := range(v) {
    if sortVals { sort.SortStrings(v) }
    out[k] = strings.Join(v, sep)
  }
  return
}

// Returns a single string created from a map[string]string, using kvsep 
// and pairsep as the seperator betwen key/value and key/value 
// pairs respectively.
func StringStringJoin(v map[string]string, kvsep string, pairsep string, sortKeys bool)(out string){
  keys := StringStringMapKeys(v)
  if sortKeys { sort.SortStrings(keys) }
  for ki := range(keys) {
    if len(out) > 0 { out += pairsep }
    out += keys[ki] + kvsep + v[keys[ki]]
  }
  return
}

// Convert a map[string]string to a map[string][]string by copying the value
// in the original map to the sole element of its respective list in the return.
func StringStringToStringStrings(in map[string]string)(out map[string][]string){
  out = make(map[string][]string)
  for k, v := range(in) {
    out[k] = []string{v}
  }
  return
}


// Return a copy of the map with escaped values.
func StringStringEscape(in map[string]string, kescape, vescape func(string)(string))(out map[string]string){
  out = map[string]string{}
  for k,v := range(in){
    if kescape != nil { k = kescape(k) }
    if vescape != nil { v = vescape(v) }
    out[k] = v
  }
  return
}
