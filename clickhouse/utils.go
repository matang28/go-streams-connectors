package clickhouse

func AsMap(record Record, mapping RecordMapping) map[string]interface{} {
	out := make(map[string]interface{})
	for idx := range record {
		out[mapping[idx]] = record[idx]
	}
	return out
}
