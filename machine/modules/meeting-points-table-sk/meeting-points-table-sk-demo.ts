import fetchMock from 'fetch-mock';

import { MeetingPointsTableSk } from './meeting-points-table-sk';

fetchMock.get('/_/meeting_points',
  [{
    PodName: 'switch-pod-0',
    Port: 12345,
    Username: 'joey-shabidoo',
    MachineID: 'skia-e-linux-123',
    LastUpdated: '2001-02-03T04:05:06.709012Z',
  },
  {
    PodName: 'switch-pod-1',
    Port: 23456,
    Username: 'chrome-boot',
    MachineID: 'skia-e-linux-234',
    LastUpdated: '2002-03-04T05:06:07.890123Z',
  }]);

const element = new MeetingPointsTableSk();
document.body.appendChild(element);
element.update();
