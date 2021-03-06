// structure
package oic

//initialize sonar struct ,call OIC initialize
func (sonar *Sonar) Init() {
	header := &sonar.Header
	OICInit(header, false)
	header.ChanNum = 3
	header.ChanOffset[0] = 0
	header.ChanOffset[1] = 11650
	header.ChanOffset[2] = 23304
	header.Channel[0].Type = 0 //sonar type:0 = sidescan 1 = angle 2 = multibeam
	header.Channel[0].Side = 0 //sonar side:0 = port 1 = starboard
	//data sample type and size:0 = 1 byte integer 1 = 2 byte integer 2 = 4 byte integer 3 = 4 byte float 4 = 12 byte set of three floats - range, theta, amp */
	header.Channel[0].Size = 1
}
