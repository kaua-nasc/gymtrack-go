import 'package:flutter/widgets.dart';

import 'app.dart';
import 'config/environment.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const GymtrackApp(environment: Environment.staging));
}
