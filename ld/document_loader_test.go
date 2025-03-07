package ld_test

import (
//"bytes"
//. "github.com/kazarena/json-gold/ld"
//"github.com/stretchr/testify/assert"
//"github.com/stretchr/testify/require"
//"testing"
)

/*
func TestLoadDocument(t *testing.T) {
	dl := NewDefaultDocumentLoader(nil)

	rd, _ := dl.LoadDocument("testdata/expand-0002-in.jsonld")

	assert.Equal(t, "t1", rd.Document.(map[string]interface{})["@type"])
}

func loadBenchData(t testing.TB) *RDFDataset {
	dl := NewDefaultDocumentLoader(nil)
	rd, err := dl.LoadDocument("testdata/compact-manifest.jsonld")
	require.Nil(t, err)
	proc := NewJsonLdProcessor()
	triples, err := proc.ToRDF(rd, NewJsonLdOptions(""))
	require.Nil(t, err)
	return triples.(*RDFDataset)
}

func BenchmarkLoadNQuads(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	err := (&NQuadRDFSerializer{}).SerializeTo(buf, loadBenchData(b))
	require.Nil(b, err)

	data := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = ParseNQuadsFrom(data)
		require.Nil(b, err)
	}
}

func TestParseLinkHeader(t *testing.T) {
	rval := ParseLinkHeader("<remote-doc-0010-context.jsonld>; rel=\"http://www.w3.org/ns/json-ld#context\"")

	assert.Equal(
		t,
		map[string][]map[string]string{
			"http://www.w3.org/ns/json-ld#context": {{
				"target": "remote-doc-0010-context.jsonld",
				"rel":    "http://www.w3.org/ns/json-ld#context",
			}},
		},
		rval,
	)
}

func TestCachingDocumentLoaderLoadDocument(t *testing.T) {
	cl := NewCachingDocumentLoader(NewDefaultDocumentLoader(nil))

	cl.PreloadWithMapping(map[string]string{
		"http://www.example.com/expand-0002-in.jsonld": "testdata/expand-0002-in.jsonld",
	})

	rd, _ := cl.LoadDocument("http://www.example.com/expand-0002-in.jsonld")

	assert.Equal(t, "t1", rd.Document.(map[string]interface{})["@type"])
}
*/
