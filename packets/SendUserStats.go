package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

// SendUserStats sends the UserStatus to the writer.
func (w *Writer) SendUserStats(x constants.UserStatsStruct) {
	p := NewPacket(constants.BanchoHandleOsuUpdate)
	p.SetPacketData(helpers.MarshalBinary(&x))
	w.Write(p.ToByteArray())
}
