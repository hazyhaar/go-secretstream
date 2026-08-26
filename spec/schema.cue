package spec

#HeaderField: {
	name:       string
	offset:     int & >=0
	size:       int & >0
	kind:       "bytes" | "u16"
	value?:     int
	value_hex?: string
}

#VectorInput: {
	name:      string
	key_hex:   =~"^[0-9a-f]{64}$"
	nonce_hex: =~"^[0-9a-f]{48}$"
	fragments_hex: [...=~"^([0-9a-f]{2})*$"]
	ads_hex: [...=~"^([0-9a-f]{2})*$"]
}

#FormatV2: {
	endianness: "big"
	header: {
		size:    int & >0
		size_v1: int & >0
		fields: [...#HeaderField]
	}
	frame: {
		length_size: int & >0
		tag_size:    int & >0
		mac_size:    int & >0
		min_payload: int & >0
		max_payload: int & >0
	}
	tags: {
		message: int & >=0 & <=255
		push:    int & >=0 & <=255
		rekey:   int & >=0 & <=255
		final:   int & >=0 & <=255
		admitted: [...int]
	}
	ad: {
		domain_size: int & >0
		seq_size:    int & >0
		tag_size:    int & >0
		len_size:    int & >0
		prefix_len:  int & >0
	}
	chunk_size: int & >0
	key_size:   int & >0
	frame_nonce: {
		prefix_from_nonce_offset: int & >=0
		prefix_size:              int & >0
		seq_size:                 int & >0
		ietf_size:                int & >0
	}
	hybrid: {
		probe_size:            int & >0
		collision_probability: string
		rules: [...string]
	}
	vectors: [...#VectorInput]
}

format_v2: #FormatV2
