/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import "github.com/deckarep/golang-set"

const VoiceServiceProfileClassId ClassID = ClassID(58)

var voiceserviceprofileBME *ManagedEntityDefinition

// VoiceServiceProfile (class ID #58)
//	This ME organizes data that describe the voice service functions of the ONU. Instances of this
//	ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of a VoIP voice CTP by way
//		of a VoIP media profile.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2 bytes)
//
//		Announcement Type
//			(R, W, setbycreate) (mandatory) (1 byte)
//
//		Jitter Target
//			Jitter target:	This attribute specifies the target value of the jitter buffer in milliseconds.
//			The system tries to maintain the jitter buffer at the target value. The value 0 specifies
//			dynamic jitter buffer sizing. (R, W, setbycreate) (optional) (2 bytes)
//
//		Jitter Buffer Max
//			Jitter buffer max: This attribute specifies the maximum depth of the jitter buffer associated
//			with this service in milliseconds. The value 0 specifies that the ONU uses its internal default.
//			(R, W, set-by-create) (optional) (2 bytes)
//
//		Echo Cancel Ind
//			Echo cancel ind: The Boolean value true specifies that echo cancellation is on; false specifies
//			off. (R, W, setbycreate) (mandatory) (1 byte)
//
//		Pstn Protocol Variant
//			PSTN protocol variant: This attribute controls which variant of POTS signalling is used on the
//			associated UNIs. Its value is equal to the [ITU-T E.164] country code. The value 0 specifies
//			that the ONU uses its internal default. (R, W, set-by-create) (optional) (2 bytes)
//
//		Dtmf Digit Levels
//			DTMF digit levels: This attribute specifies the power level of DTMF digits that may be generated
//			by the ONU towards the subscriber set. It is a 2s complement value referred to 1 mW at the 0
//			transmission level point (TLP) (dBm0), with resolution 1 dB. The default value 0x8000 selects
//			the ONU's internal policy. (R, W, setbycreate) (optional) (2 bytes)
//
//		Dtmf Digit Duration
//			DTMF digit duration: This attribute specifies the duration of DTMF digits that may be generated
//			by the ONU towards the subscriber set. It is specified in milliseconds. The default value 0
//			selects the ONU's internal policy. (R, W, setbycreate) (optional) (2 bytes)
//
//		Hook Flash Minimum Time
//			Hook flash minimum time: This attribute defines the minimum duration recognized by the ONU as a
//			switchhook flash. It is expressed in milliseconds; the default value 0 selects the ONU's
//			internal policy. (R, W, setbycreate) (optional) (2 bytes)
//
//		Hook Flash Maximum Time
//			Hook flash maximum time: This attribute defines the maximum duration recognized by the ONU as a
//			switchhook flash. It is expressed in milliseconds; the default value 0 selects the ONU's
//			internal policy. (R, W, setbycreate) (optional) (2 bytes)
//
//		Tone Pattern Table
//			(R, W) (optional) (N * 20 bytes)
//
//		Tone Event Table
//			(R, W) (optional) (N * 7 bytes).
//
//		Ringing Pattern Table
//			(R, W) (optional) (N * 5 bytes).
//
//		Ringing Event Table
//			(R, W) (optional) (N * 7 bytes).
//
//		Network Specific Extensions Pointer
//			Network specific extensions pointer: This attribute points to a network address ME that contains
//			the path and name of a file containing network specific parameters for the associated UNIs. The
//			default value for this attribute is 0xFFFF, a null pointer. (R, W, set-by-create) (optional)
//			(2 bytes)
//
type VoiceServiceProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	voiceserviceprofileBME = &ManagedEntityDefinition{
		Name:    "VoiceServiceProfile",
		ClassID: 58,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFC,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  ByteField("AnnouncementType", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2:  Uint16Field("JitterTarget", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 2),
			3:  Uint16Field("JitterBufferMax", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 3),
			4:  ByteField("EchoCancelInd", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
			5:  Uint16Field("PstnProtocolVariant", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 5),
			6:  Uint16Field("DtmfDigitLevels", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 6),
			7:  Uint16Field("DtmfDigitDuration", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 7),
			8:  Uint16Field("HookFlashMinimumTime", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 8),
			9:  Uint16Field("HookFlashMaximumTime", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 9),
			10: MultiByteField("TonePatternTable", 20, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 10),
			11: MultiByteField("ToneEventTable", 7, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 11),
			12: MultiByteField("RingingPatternTable", 5, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 12),
			13: MultiByteField("RingingEventTable", 7, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 13),
			14: Uint16Field("NetworkSpecificExtensionsPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 14),
		},
	}
}

// NewVoiceServiceProfile (class ID 58 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVoiceServiceProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(voiceserviceprofileBME, params...)
}
