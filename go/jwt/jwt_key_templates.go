// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package jwt

import (
	"google.golang.org/protobuf/proto"
	jwtmacpb "github.com/google/tink/go/proto/jwt_hmac_go_proto"
	tinkpb "github.com/google/tink/go/proto/tink_go_proto"
)

func createJWTHMACKeyTemplate(keySize uint32, algorithm jwtmacpb.JwtHmacAlgorithm, outputPrefixType tinkpb.OutputPrefixType) *tinkpb.KeyTemplate {
	format := &jwtmacpb.JwtHmacKeyFormat{
		KeySize:   keySize,
		Version:   jwtHMACKeyVersion,
		Algorithm: algorithm,
	}
	serializedFormat, _ := proto.Marshal(format)
	return &tinkpb.KeyTemplate{
		TypeUrl:          jwtHMACTypeURL,
		Value:            serializedFormat,
		OutputPrefixType: outputPrefixType,
	}
}

// HS256Template creates a JWT key template for JWA algorithm "HS256", which is a HMAC-SHA256 with a 32 byte key. It will set a key ID header "kid" in the token.
func HS256Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(32, jwtmacpb.JwtHmacAlgorithm_HS256, tinkpb.OutputPrefixType_TINK)
}

// RawHS256Template creates a JWT key template for JWA algorithm "HS256", which is a HMAC-SHA256 with a 32 byte key. It will not set a key ID header "kid" in the token.
func RawHS256Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(32, jwtmacpb.JwtHmacAlgorithm_HS256, tinkpb.OutputPrefixType_RAW)
}

// HS384Template creates a JWT key template for JWA algorithm "HS384", which is a HMAC-SHA384 with a 48 byte key. It will set a key ID header "kid" in the token.
func HS384Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(48, jwtmacpb.JwtHmacAlgorithm_HS384, tinkpb.OutputPrefixType_TINK)
}

// RawHS384Template creates a JWT key template for JWA algorithm "HS384", which is a HMAC-SHA384 with a 48 byte key. It will not set a key ID header "kid" in the token.
func RawHS384Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(48, jwtmacpb.JwtHmacAlgorithm_HS384, tinkpb.OutputPrefixType_RAW)
}

// HS512Template creates a JWT key template for JWA algorithm "HS512", which is a HMAC-SHA512 with a 64 byte key. It will set a key ID header "kid" in the token.
func HS512Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(64, jwtmacpb.JwtHmacAlgorithm_HS512, tinkpb.OutputPrefixType_TINK)
}

// RawHS512Template creates a JWT key template for JWA algorithm "HS512", which is a HMAC-SHA512 with a 64 byte key. It will not set a key ID header "kid" in the token.
func RawHS512Template() *tinkpb.KeyTemplate {
	return createJWTHMACKeyTemplate(64, jwtmacpb.JwtHmacAlgorithm_HS512, tinkpb.OutputPrefixType_RAW)
}
