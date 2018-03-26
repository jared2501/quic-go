package congestion

import (
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"time"
)

type FixedCongestionWindowSender struct {
	CongestionWindow protocol.ByteCount
	RttStats         *RTTStats
}

func (s *FixedCongestionWindowSender) TimeUntilSend(bytesInFlight protocol.ByteCount) time.Duration {
	return s.RttStats.SmoothedRTT() / time.Duration(2*s.CongestionWindow/protocol.DefaultTCPMSS)
}

func (s *FixedCongestionWindowSender) OnPacketSent(sentTime time.Time, bytesInFlight protocol.ByteCount, packetNumber protocol.PacketNumber, bytes protocol.ByteCount, isRetransmittable bool) bool {
	return isRetransmittable
}

func (s *FixedCongestionWindowSender) GetCongestionWindow() protocol.ByteCount {
	return s.CongestionWindow
}

func (s *FixedCongestionWindowSender) MaybeExitSlowStart() {
}

func (s *FixedCongestionWindowSender) OnPacketAcked(number protocol.PacketNumber, ackedBytes protocol.ByteCount, bytesInFlight protocol.ByteCount) {
}

func (s *FixedCongestionWindowSender) OnPacketLost(number protocol.PacketNumber, lostBytes protocol.ByteCount, bytesInFlight protocol.ByteCount) {
}

func (s *FixedCongestionWindowSender) SetNumEmulatedConnections(n int) {
}

func (s *FixedCongestionWindowSender) OnRetransmissionTimeout(packetsRetransmitted bool) {
}

func (s *FixedCongestionWindowSender) OnConnectionMigration() {
}

func (s *FixedCongestionWindowSender) RetransmissionDelay() time.Duration {
	if s.RttStats.SmoothedRTT() == 0 {
		return 0
	}
	return s.RttStats.SmoothedRTT() + s.RttStats.MeanDeviation()*4
}

func (s *FixedCongestionWindowSender) SetSlowStartLargeReduction(enabled bool) {
}
