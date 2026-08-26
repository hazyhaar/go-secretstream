package spec

format_v2: {
	let cs = 65536
	let mac = 16
	let mag = 8
	let seqb = 8
	let tagb = 1
	let adlen = 4

	endianness: "big"

	header: {
		size:    36
		size_v1: 24
		fields: [
			{name: "magic", offset: 0, size: 8, kind: "bytes", value_hex: "535335352d763200"},
			{name: "version", offset: 8, size: 2, kind: "u16", value: 2},
			{name: "flags", offset: 10, size: 2, kind: "u16", value: 0},
			{name: "nonce", offset: 12, size: 24, kind: "bytes"},
		]
		size: fields[3].offset + fields[3].size
	}

	frame: {
		length_size: 4
		tag_size:    tagb
		mac_size:    mac
		min_payload: tag_size + mac_size
		max_payload: tag_size + cs + mac_size
	}

	tags: {
		message: 0x00
		push:    0x01
		rekey:   0x02
		final:   0x03
		admitted: [message, final]
	}

	ad: {
		domain_size: mag
		seq_size:    seqb
		tag_size:    tagb
		len_size:    adlen
		prefix_len:  mag + seqb + tagb + adlen
	}

	chunk_size: cs
	key_size:   32

	frame_nonce: {
		prefix_from_nonce_offset: 16
		prefix_size:              4
		seq_size:                 8
		ietf_size:                prefix_size + seq_size
	}

	hybrid: {
		probe_size:            8
		collision_probability: "2^-64"
		rules: [
			"Si les huit premiers octets valent le magique, le reste de l'en-tête v2 est consommé.",
			"Sinon ces huit octets sont le début du nonce v1 et les seize suivants le complètent.",
			"Le format v1 n'a ni magique, ni version, ni bloc terminal. Une archive v1 qui s'arrête sur une frontière de trame se lit comme complète.",
		]
	}

	vectors: [
		{
			name:      "empty_close"
			key_hex:   "515c4b7665101f0a3924d3decdf8e792818cbba655404f7a6914030e3d28d7c2"
			nonce_hex: "c3c4cdd6dfe0e9f2fbfc858e9798a1aab3b4bd464f505962"
			fragments_hex: []
			ads_hex: []
		},
		{
			name:      "hello_no_ad"
			key_hex:   "515c4b7665101f0a3924d3decdf8e792818cbba655404f7a6914030e3d28d7c2"
			nonce_hex: "c3c4cdd6dfe0e9f2fbfc858e9798a1aab3b4bd464f505962"
			fragments_hex: ["68656c6c6f"]
			ads_hex: [""]
		},
		{
			name:      "msg_with_ad"
			key_hex:   "515c4b7665101f0a3924d3decdf8e792818cbba655404f7a6914030e3d28d7c2"
			nonce_hex: "c3c4cdd6dfe0e9f2fbfc858e9798a1aab3b4bd464f505962"
			fragments_hex: ["032c557ea7d0f9224b749dc6ef18416a93bce50e376089b2db042d567fa8d1fa234c759ec7f019426b94bde60f38618ab3dc052e5780a9d2fb244d769fc8f11a"]
			ads_hex: ["626f756e64"]
		},
		{
			name:      "two_frags_distinct_ad"
			key_hex:   "515c4b7665101f0a3924d3decdf8e792818cbba655404f7a6914030e3d28d7c2"
			nonce_hex: "c3c4cdd6dfe0e9f2fbfc858e9798a1aab3b4bd464f505962"
			fragments_hex: ["616c706861", "62657461"]
			ads_hex: ["78", "79"]
		},
	]
}
